package services

import (
	"sync"
	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
	"fmt"
	"bytes"
	"github.com/matiasdominguez/pass/src/api/logger"
	"image/png"
	"github.com/lazywei/go-opencv/opencv"
)

var instance *openalpr.Alpr
var once sync.Once

func GetInstance() *openalpr.Alpr {
	once.Do(func() {
		instance = openalpr.NewAlpr("ar", "/usr/local/Cellar/openalpr/HEAD-3a86c5f_1/share/openalpr/runtime_data/config/ar.conf", "/usr/local/Cellar/openalpr/HEAD-3a86c5f_1/share/openalpr/runtime_data/")
		instance.SetTopN(2)
		instance.SetDefaultRegion("ar")
	})
	return instance
}

func AlprStream() {
	cap := opencv.NewCameraCapture(0)
	cap.SetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH, 300)
	cap.SetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT, 300)
	defer cap.Release()
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			buff := new(bytes.Buffer)
			err := png.Encode(buff, img.ToImage())
			if err != nil {
				fmt.Println(err)
			}
			Recognize(buff)
		}
	}
}

func Recognize(buff *bytes.Buffer) {
	// TODO REMOVE HARDCODING
	go CreateEventIfCarExists("FMU526")
	resultFromFilePath, err := GetInstance().RecognizeByBlob(buff.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	if len(resultFromFilePath.Plates) > 0 {
		//logger.Info(fmt.Sprintf("%+v\n", resultFromFilePath.Plates[0].BestPlate))
		logger.Info(resultFromFilePath.Plates[0].BestPlate)
	}
}