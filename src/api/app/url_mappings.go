package app

import (
	"github.com/matiasdominguez/pass/src/api/controllers"
)

func mapUrlsToControllers() {
	//Router.POST("/bigqueue/freepass/process_payment", fpController.ProcessMessage)
	//Router.GET("/webcam/stream", controllers.StreamWebcam)
	Router.POST("/resident", controllers.CreateResident)
	Router.GET("/residents", controllers.ListResidents)
	Router.GET("/cars", controllers.ListCars)
}