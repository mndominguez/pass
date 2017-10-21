package services

import (
	"github.com/matiasdominguez/pass/src/api/domain"
	"fmt"
	"time"
	"github.com/matiasdominguez/pass/src/api/logger"
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
}

func InsertEvent(event *domain.Event) error {
	db := GetDBInstance()
	err := db.Insert(event)
	if err != nil {
		return err
	}
	return nil
}

func GetResident(car domain.Car) (domain.Resident, error) {
	db := GetDBInstance()
	resident := domain.Resident{Cars: []domain.Car{car}}
	err := db.Select(&resident)
	if err != nil {
		return resident, err
	}
	fmt.Print(resident)
	return resident, nil
}