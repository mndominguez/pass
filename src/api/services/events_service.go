package services

import (
	"github.com/matiasdominguez/pass/src/api/domain"
	"github.com/matiasdominguez/pass/src/api/logger"
	"fmt"
	"time"
)

func CreateEventIfCarExists(plate string) {
	car, err := GetCar(plate)
	if err != nil {
		// Car is not on database
		// do nothing
	}
	if err == nil {
		// Get owner
		fmt.Println(car.Residents[0].Id)
		// Create event and insert it
		event := &domain.Event{
			Car: &car,
			Resident: &car.Residents[0],
			CreationDate: time.Now(),
		}
		err := InsertEvent(event)
		if err != nil {
			panic(err)
		}
		logger.Info(fmt.Sprintf("%+v\n", event))
	}
	go FindEventsByCar(domain.Car{Plate:"FMU526"})
}

func InsertEvent(event *domain.Event) error {
	db := GetDBInstance()
	err := db.Insert(event)
	if err != nil {
		return err
	}
	return nil
}

func FindEventsByCar(car domain.Car) {
	db := GetDBInstance()
	var events []domain.Event
	err := db.Model(&events).Where("car ->>'Plate' = ?", car.Plate).Select()
	if err != nil {
		panic(err)
	}
	logger.Info(fmt.Sprintf("%+v\n", events[0].Car.Plate))
	logger.Info(fmt.Sprintf("%+v\n", events))
}