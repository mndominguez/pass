package old

import (
	"fmt"
//	"os"
	//"path"
	//"runtime"

	"github.com/lazywei/go-opencv/opencv"
	//"../opencv" // can be used in forks, comment in real application
	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
//	"bytes"
	//"image/jpeg"
//	"image/png"
//	"bufio"
	"bytes"
	"image/png"
//	"io"
)

func main() {
	alpr := openalpr.NewAlpr("ar", "/usr/local/Cellar/openalpr/HEAD-3a86c5f_1/share/openalpr/runtime_data/config/ar.conf", "/usr/local/Cellar/openalpr/HEAD-3a86c5f_1/share/openalpr/runtime_data/")
	defer alpr.Unload()

	if !alpr.IsLoaded() {
		fmt.Println("OpenAlpr failed to load!")
		return
	}
	alpr.SetTopN(5)

	alpr.SetDefaultRegion("ar")

	fmt.Println(alpr.IsLoaded())
	fmt.Println(openalpr.GetVersion())

	//mainVideo(alpr)
	newWay(alpr)
}

func newWay(alpr *openalpr.Alpr) {
	cap := opencv.NewCameraCapture(0)
	//cap.SetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH, 640)
	//cap.SetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT, 480)
	defer cap.Release()
	limit := 0
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if limit == 1 {
				limit = 0
				getPrediction(img, alpr)
			}
		}
		limit += 1
	}
}

func mainVideo(alpr *openalpr.Alpr) *opencv.IplImage {
	win := opencv.NewWindow("Go-OpenCV Webcam")
	defer win.Destroy()
	cap := opencv.NewCameraCapture(0)
	cap.SetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH, 640)
	cap.SetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT, 480)
	if cap == nil {
		panic("can not open camera")
	}
	defer cap.Release()
	fmt.Println("Press ESC to quit")
	limit := 0
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if img != nil {
				//win.ShowImage(img)
				if limit == 5 {
					limit = 0
					go getPrediction(img, alpr)
				}
			} else {
				fmt.Println("Image ins nil")
			}
		}
		opencv.WaitKey(1)

		//if key == 27 {
		//	os.Exit(0)
		//}
		limit += 1
	}
}

func getPrediction(img *opencv.IplImage, alpr *openalpr.Alpr) {
	buff := new(bytes.Buffer)
	err := png.Encode(buff, img.ToImage())
	resultFromFilePath, err := alpr.RecognizeByBlob(buff.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resultFromFilePath.Plates)
}
