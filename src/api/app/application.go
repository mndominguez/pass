package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/matiasdominguez/pass/src/api/logger"
	//"github.com/matiasdominguez/pass/src/api/services"
	"github.com/matiasdominguez/pass/src/api/services"
)

var (
	Router *gin.Engine
)


func StartApp() {
	services.StartModel()
	services.AlprStream()
	configureRouter()
	mapUrlsToControllers()
	//services.StartModel()
	Router.Run(":8080")
}

func configureRouter() {
	Router = gin.Default()
	gin.SetMode(gin.DebugMode)

	// Creates the gin engine
	Router.RedirectFixedPath = false
	Router.RedirectTrailingSlash = false
	Router.HandleMethodNotAllowed = true

	log.Info("Application successfully configured.")
}
