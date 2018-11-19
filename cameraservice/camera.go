package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

//Testing GOCV library.
//Capture a single frame from webcam, then save it to an image file on disk.
//Directory name => ../lib/photoData

//Task 1 => Read documentation on GOCV
//Task 2 => Capture video from webcam, then detect face from user.
//Task 3 => Send pictue to database. MongoDB
func main() {
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		panic(err)
	}
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Println("Cannot read device")
		return
	}
	if img.Empty() {
		fmt.Println("no image on device")
		return
	}
	window.IMShow(img)
	window.WaitKey(1)
	gocv.IMWrite("../lib/photoData/filename.jpg", img)
}
