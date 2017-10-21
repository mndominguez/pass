package services

import (
	"github.com/matiasdominguez/pass/src/api/domain"
	//"fmt"
)

func GetCar(plate string) (domain.Car, error) {
	db := GetDBInstance()
	car := domain.Car{Plate: plate}
	err := db.Select(&car)
	if err != nil {
		return car, err
	}
	//fmt.Print(car)
	return car, nil
}