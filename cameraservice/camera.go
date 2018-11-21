package main

import (
	"fmt"
	"image"
	"image/color"
	"path/filepath"

	"gocv.io/x/gocv"
)

//Testing GOCV library.
//Capture a single frame from webcam, then save it to an image file on disk.
//Directory name => ../lib/photoData

//Completed. Task 1 => Read documentation on GOCV
// Completed. Task 2 => Capture video from webcam, then detect face from user. Using a counter GOCV will take a picture right when counter is 10 seconds and when blob is draw.
//Task 3 => Send pictue to database. MongoDB
var model string
var config string

func main() {
	backend := gocv.NetBackendDefault
	target := gocv.NetTargetCPU
	model = "res10_300x300_ssd_iter_140000.caffemodel"
	config = "deploy.prototxt"
	window := gocv.NewWindow("FaceID")

	net := gocv.ReadNet(model, config)
	if net.Empty() {
		err := fmt.Errorf("Error reading network model")
		panic(err)
	}

	net.SetPreferableBackend(backend)
	net.SetPreferableTarget(target)

	err := setVideo(0, window, net)
	if err != nil {
		panic(err)
	}
	defer net.Close()
}

func setVideo(inputID int, window *gocv.Window, net gocv.Net) error {
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

		performDetection(&img, prob)
		window.IMShow(img)
		window.WaitKey(1)
	}

}

// performDetection analyzes the results from the detector network,
// which produces an output blob with a shape 1x1xNx7
// where N is the number of detections, and each detection
// is a vector of float values
// [batchId, classId, confidence, left, top, right, bottom]
func performDetection(frame *gocv.Mat, results gocv.Mat) {
	for i := 0; i < results.Total(); i += 7 {
		confidence := results.GetFloatAt(0, i+2)
		if confidence > 0.5 {
			left := int(results.GetFloatAt(0, i+3) * float32(frame.Cols()))
			top := int(results.GetFloatAt(0, i+4) * float32(frame.Rows()))
			right := int(results.GetFloatAt(0, i+5) * float32(frame.Cols()))
			bottom := int(results.GetFloatAt(0, i+6) * float32(frame.Rows()))
			gocv.Rectangle(frame, image.Rect(left, top, right, bottom), color.RGBA{0, 255, 0, 0}, 2)
		}
	}
}
