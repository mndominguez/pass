package controllers

import (
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"github.com/matiasdominguez/pass/src/api/logger"
	"net/http"
	//"github.com/matiasdominguez/pass/src/api/domain"
	//"encoding/json"
	//"fmt"
	"github.com/matiasdominguez/pass/src/api/services"
	"io/ioutil"
	"fmt"
	"github.com/matiasdominguez/pass/src/api/domain"
	"encoding/json"
)

type Message struct {
	Message  string `json:"message"`
	ErrorStr string `json:"error"`
	Status   int    `json:"status"`
	Cause    string `json:"cause"`
}

func CreateResident(c *gin.Context) {
	req, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("Error trying to parse json body", err)
		c.JSON(http.StatusBadRequest, Message{Message: "invalid json body", Status: http.StatusBadRequest})
		return
	}
	var resident domain.Resident
	err = json.Unmarshal(req, &resident)
	var car domain.Car
	err = json.Unmarshal(req, &car)
	logger.Info(fmt.Sprintf("Got resident: %v", resident))
	logger.Info(fmt.Sprintf("Got car: %v", car))
	if err != nil {
		logger.Error("Error trying to parse json body", err)
		c.JSON(http.StatusBadRequest, Message{Message: "invalid json body", Status: http.StatusBadRequest})
		return
	}
	db := services.GetDBInstance()
	err = db.Insert(&resident)
	if err != nil {
		logger.Error("Error trying to insert resident", err)
		c.JSON(http.StatusBadRequest, Message{Message: "couldn't insert resident", Status: http.StatusBadRequest})
		return
	}
	car.ResidentId = resident.Id
	err = db.Insert(&car)
	if err != nil {
		logger.Error("Error trying to insert car", err)
		c.JSON(http.StatusBadRequest, Message{Message: "couldn't insert car", Status: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, Message{Message:"OK", Status: http.StatusOK})
}

func ListResidents(c *gin.Context) {
	dbresidents, err := services.ListAllResidents()
	if err != nil {
		logger.Error("Error accesing database", err)
		c.JSON(http.StatusInternalServerError, Message{Message: "Couldn't get that list", Status: http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, dbresidents)
}

func ListCars(c *gin.Context) {
	dbcars, err := services.ListCars()
	if err != nil {
		logger.Error("Error accesing database", err)
		c.JSON(http.StatusInternalServerError, Message{Message: "Couldn't get that list", Status: http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, dbcars)
}

func GetEvents(c *gin.Context) {

}

func GetResident(c *gin.Context) {
	name := c.Query("name")
	dbresident, err := services.GetResidentByName(name)
	if err != nil {
		logger.Error("Error trying to parse json body", err)
		c.JSON(http.StatusNotFound, Message{Message: "There's no one by that name", Status: http.StatusNotFound})
		return
	}
	c.JSON(http.StatusOK, dbresident)
}