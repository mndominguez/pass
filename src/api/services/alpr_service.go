package services

import (
	"sync"
	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
	"fmt"
	"bytes"
	"github.com/mndominguez/pass/src/api/logger"
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

func Recognize(buff *bytes.Buffer) {
	resultFromFilePath, err := GetInstance().RecognizeByBlob(buff.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	if len(resultFromFilePath.Plates) > 0 {
		//logger.Info(fmt.Sprintf("%+v\n", resultFromFilePath.Plates[0].BestPlate))
		logger.Info(resultFromFilePath.Plates[0].BestPlate)
	}
}