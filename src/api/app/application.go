package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/mndominguez/pass/src/api/logger"
	"github.com/mndominguez/pass/src/api/services"
)

var (
	Router *gin.Engine
)


func StartApp() {
	configureRouter()
	mapUrlsToControllers()
	services.StartModel()
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
