package old

import (
	"fmt"
	"io/ioutil"
)

import "github.com/openalpr/openalpr/src/bindings/go/openalpr"

func main() {
	alpr := openalpr.NewAlpr("ar", "/usr/local/Cellar/openalpr/HEAD-3a86c5f_1/share/openalpr/runtime_data/config/ar.conf", "/usr/local/Cellar/openalpr/HEAD-3a86c5f_1/share/openalpr/runtime_data/")
	defer alpr.Unload()

	if !alpr.IsLoaded() {
		fmt.Println("OpenAlpr failed to load!")
		return
	}
	alpr.SetTopN(20)

	alpr.SetDefaultRegion("ar")

	fmt.Println(alpr.IsLoaded())
	fmt.Println(openalpr.GetVersion())

	resultFromFilePath, err := alpr.RecognizeByFilePath("/Users/mdominguez/Documents/Developer/mko_alpr_qt/test_images/FMU526.jpg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resultFromFilePath)
	fmt.Printf("\n\n\n")

	imageBytes, err := ioutil.ReadFile("/Users/mdominguez/Documents/Developer/mko_alpr_qt/test_images/FMU526.jpg")
	if err != nil {
		fmt.Println(err)
	}
	resultFromBlob, err := alpr.RecognizeByBlob(imageBytes)
	fmt.Printf("%+v\n", resultFromBlob)
}
