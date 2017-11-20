package services

import (
	"fmt"

	"github.com/matiasdominguez/pass/src/api/domain"
	"github.com/go-pg/pg/orm"
)

func InsertNewResident(resident domain.Resident) error {
	db := GetDBInstance()
	err := db.Insert(resident)
	if err != nil {
		return err
	}
	return nil
}

func GetResidentByName(name string) (domain.Resident, error) {
	db := GetDBInstance()
	resident := domain.Resident{Name: name}
	err := db.Select(&resident)
	if err != nil {
		return resident, err
	}
	fmt.Println(resident)
	return resident, nil
}

func ListAllResidents() ([]domain.Resident, error) {
	db := GetDBInstance()
	var residents []domain.Resident
	//err := db.Model(&residents).Order("id asc").Select()
	err := db.Model(&residents).
		Column("resident.*", "Cars").
		Relation("Cars", func(q *orm.Query) (*orm.Query, error) {
		return q, nil
		}).
		Select()
	fmt.Println(err)
	fmt.Println(residents)
	return residents, nil
}

/*
func GetResidentByCar(car domain.Car) (domain.Resident, error) {
	db := GetDBInstance()
	resident := domain.Resident{Cars: []domain.Car{car}}
	err := db.Select(&resident)
	if err != nil {
		return resident, err
	}
	fmt.Print(resident)
	return resident, nil
}*/
