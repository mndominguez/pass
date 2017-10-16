package app

import (
	webcamController "github.com/matiasdominguez/pass/src/api/controllers"
)

func mapUrlsToControllers() {
	//Router.POST("/bigqueue/freepass/process_payment", fpController.ProcessMessage)
	Router.GET("/webcam/stream", webcamController.StreamWebcam)
}