package old

import (
//	"fmt"
	"log"
	"net/http"
	"os/exec"
	"github.com/lazywei/go-opencv/opencv"
	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
	"bytes"
	"image/png"
	"fmt"
)

const boundary = "informs"

func main() {
	http.HandleFunc("/cam", cam)
	http.HandleFunc("/camFacebox", camFacebox)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func camFacebox(w http.ResponseWriter, r *http.Request) {
	alpr := openalpr.NewAlpr("ar", "/usr/local/Cellar/openalpr/HEAD-3a86c5f_1/share/openalpr/runtime_data/config/ar.conf", "/usr/local/Cellar/openalpr/HEAD-3a86c5f_1/share/openalpr/runtime_data/")
	alpr.SetTopN(2)
	alpr.SetDefaultRegion("ar")
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary="+boundary)
	cap := opencv.NewCameraCapture(0)
	cap.SetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH, 300)
	cap.SetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT, 300)
	defer cap.Release()
	limit := 0
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			buff := new(bytes.Buffer)
			err := png.Encode(buff, img.ToImage())
			reader := bytes.NewReader(buff.Bytes())
			if err != nil {
				fmt.Println(err)
			}
			w.Write([]byte("Content-Type: image/jpeg\r\n"))
			w.Write([]byte("Content-Length: " + string(reader.Len()) + "\r\n\r\n"))
			w.Write(buff.Bytes())
			w.Write([]byte("\r\n"))
			w.Write([]byte("--informs\r\n"))
			resultFromFilePath, err := alpr.RecognizeByBlob(buff.Bytes())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%+v\n", resultFromFilePath.Plates)
		}
		limit += 1
	}
}
