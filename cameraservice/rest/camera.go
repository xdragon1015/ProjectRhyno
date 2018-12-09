package rest

import (
	"ProjectRhyno/lib/persistance"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gocv.io/x/gocv"
)

//Testing GOCV library.
//Capture a single frame from webcam, then save it to an image file on disk.
//Directory name => ../lib/photoData

//Completed. Task 1 => Read documentation on GOCV
// Completed. Task 2 => Capture video from webcam, then detect face from user. Using a counterer GOCV will take a picture right when counter is 10 seconds and when blob is draw.
//Task 3 => Send pictue to database. MongoDB

const (
	windowName     = "FaceID"
	modeName       = "res10_300x300_ssd_iter_140000.caffemodel"
	configFileName = "deploy.prototxt"
)

var (
	numberOfDetections = 0
	counter            = 0
)

type CameraServiceHandler struct {
	dbhandler persistance.DatabaseHandler
}

//NewCameraServiceHandler initializises a CameraServiceHandler object with dbHandler for which allows to work directly with the database
func NewCameraServiceHandler(databaseHandler persistance.DatabaseHandler) *CameraServiceHandler {
	return &CameraServiceHandler{
		dbhandler: databaseHandler,
	}
}

var model string
var config string

//SetCameraService creates window, net, and sets the target and preferable backend to the net from the gocv package
func SetCameraService() {
	backend := gocv.NetBackendDefault
	target := gocv.NetTargetCPU
	model = modeName
	config = configFileName
	window := gocv.NewWindow(windowName)

	net := gocv.ReadNet(model, config)
	if net.Empty() {
		err := fmt.Errorf("Error reading network model")
		panic(err)
	}

	net.SetPreferableBackend(backend)
	net.SetPreferableTarget(target)

	err := SetVideo(0, window, net)
	if err != nil {
		panic(err)
	}
	defer net.Close()
}

//SetVideo initializes the video recording and face detection features
func SetVideo(inputID int, window *gocv.Window, net gocv.Net) error {
	var ratio float64
	var mean gocv.Scalar
	var swapRGB bool
	webcam, err := gocv.OpenVideoCapture(inputID)
	if err != nil {
		return err
	}

	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			var err = fmt.Errorf("Cannot read device: Return bool is %v", ok)
			return err
		}
		if img.Empty() {
			err := fmt.Errorf("No image on device")
			return err
		}

		if filepath.Ext(model) == ".caffemodel" {
			ratio = 1.0
			mean = gocv.NewScalar(4, 127, 120, 0)
			swapRGB = false
		} else {
			ratio = 1.0 / 127.5
			mean = gocv.NewScalar(127.5, 127.5, 127.5, 0)
			swapRGB = true
		}
		blob := gocv.BlobFromImage(img, ratio, image.Pt(300, 300), mean, swapRGB, false)

		// feed the blob into the detector
		net.SetInput(blob, "")

		prob := net.Forward("")
		//Perform detections until detections equals 0
		if detections := PerformDetection(&img, prob); detections != 0 {
			time.Sleep(time.Microsecond * 10000)
			TakePhoto(&img, counter)
			//Update counter accorsindly to the number of detections thats been detected
			counter++
			//Set numberOdDetections back to 0
			numberOfDetections = 0
		}

		window.IMShow(img)
		window.WaitKey(1)
	}
	return nil
}

// PerformDetection analyzes the results from the detector network,
// which produces an output blob with a shape 1x1xNx7
// where N is the number of detections, and each detection
// is a vector of float values
// [batchId, classId, confidence, left, top, right, bottom]
func PerformDetection(frame *gocv.Mat, results gocv.Mat) int {

	for i := 0; i < results.Total(); i += 7 {
		confidence := results.GetFloatAt(0, i+2)
		if confidence > 0.75 {
			left := int(results.GetFloatAt(0, i+3) * float32(frame.Cols()))
			top := int(results.GetFloatAt(0, i+4) * float32(frame.Rows()))
			right := int(results.GetFloatAt(0, i+5) * float32(frame.Cols()))
			bottom := int(results.GetFloatAt(0, i+6) * float32(frame.Rows()))
			gocv.Rectangle(frame, image.Rect(left, top, right, bottom), color.RGBA{0, 255, 0, 0}, 2)
			numberOfDetections++
			// takePhoto(frame, numberOfDetections)
		}
	}
	if numberOfDetections > 30 {
		return numberOfDetections
	}
	fmt.Println(numberOfDetections)
	return 0
}

func getPictures() []byte {
	var photo []byte
	filepath.Walk("photos", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		file, err := ioutil.ReadFile(path)
		photo = file
		return err
	})
	return photo
}

//TakePhoto does not overrides pictures after updating to n=counter
func TakePhoto(frame *gocv.Mat, n int) {
	url := fmt.Sprintf("../lib/photos/photoData%v.jpg", n)
	fmt.Println(gocv.IMWrite(url, *frame))
}

func (eh *CameraServiceHandler) addPictureHandler(w http.ResponseWriter, r *http.Request) {
	pic := persistance.Photo{}
	err := json.NewDecoder(r.Body).Decode(&pic)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occurred while persisting picture %s}", err)
		return
	}

	id, err := eh.dbhandler.AddPhoto(pic)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occurred while persisting picture %d %s}", id, err)
		return
	}
}
