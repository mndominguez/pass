package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/matiasdominguez/pass/src/api/logger"
	"net/http"
	"github.com/matiasdominguez/pass/src/api/domain"
	"encoding/json"
	"fmt"
)

type ApiError struct {
	Message  string `json:"message"`
	ErrorStr string `json:"error"`
	Status   int    `json:"status"`
	Cause    string `json:"cause"`
}

func CreateResident(c *gin.Context) {
	req, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		logger.Error("Error trying to parse json body", err)
		c.JSON(http.StatusBadRequest, ApiError{Message: "invalid json body", Status: http.StatusBadRequest})
		return
	}

	var resident domain.Resident

	err = json.Unmarshal(req, &resident)

	var car domain.Car

	json.Unmarshal([]byte(req), &n)

	logger.Info(fmt.Sprintf("Got message: %v", resident))
	if err != nil {
		logger.Error("Error trying to parse json body", err)
		c.JSON(http.StatusBadRequest, ApiError{Message: "invalid json body", Status: http.StatusBadRequest})
		return
	}
}