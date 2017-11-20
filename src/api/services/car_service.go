package services

import (
	"github.com/matiasdominguez/pass/src/api/domain"
	//"fmt"
	"fmt"
//	"github.com/go-pg/pg/orm"
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

func ListCars() ([]domain.Car, error) {
	db := GetDBInstance()
	var cars []domain.Car
	err := db.Model(&cars).Select()
	fmt.Println(err)
	fmt.Println(cars)
	return cars, nil
}