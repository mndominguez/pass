package controllers

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"image/png"
	"fmt"
	"github.com/lazywei/go-opencv/opencv"
//	"github.com/matiasdominguez/pass/src/api/services"
)

const boundary = "informs"

func StreamWebcam(c *gin.Context) {
	c.Header("Content-Type", "multipart/x-mixed-replace; boundary="+boundary)
	cap := opencv.NewCameraCapture(0)
	cap.SetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH, 300)
	cap.SetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT, 300)
	defer cap.Release()
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			buff := new(bytes.Buffer)
			err := png.Encode(buff, img.ToImage())
			reader := bytes.NewReader(buff.Bytes())
			if err != nil {
				fmt.Println(err)
			}
			c.Writer.Write([]byte("Content-Type: image/jpeg\r\n"))
			c.Writer.Write([]byte("Connection: keep-alive\r\n"))
			c.Writer.Write([]byte("Content-Length: " + string(reader.Len()) + "\r\n\r\n"))
			c.Writer.Write(buff.Bytes())
			c.Writer.Write([]byte("\r\n"))
			c.Writer.Write([]byte("--informs\r\n"))
			//services.Recognize(buff)
		}
	}
}
