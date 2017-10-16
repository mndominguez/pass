package services

import (
	"github.com/matiasdominguez/pass/src/api/domain"
)

func GetCar(plate string) domain.Car {
	db := GetDBInstance()
	car := domain.Car{Plate: plate}
	err := db.Select(&car)
	if err != nil {
		panic(err)
	}
	return car
}