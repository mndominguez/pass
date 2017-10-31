package services

import (
	"fmt"

	"github.com/matiasdominguez/pass/src/api/domain"
)

func InsertNewResident(resident domain.Resident) error {
	db := GetDBInstance()
	err := db.Insert(resident)
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
