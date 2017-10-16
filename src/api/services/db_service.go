package services

import (
	"github.com/go-pg/pg"
	"github.com/matiasdominguez/pass/src/api/domain"
)

func StartModel() {
	db := pg.Connect(&pg.Options{
		User: "mdominguez",
		Database: "pass",
	})

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	user1 := &domain.User{
		Username: "admin",
		Password: "admin",
		Role: "admin",
	}
	err = db.Insert(user1)
	if err != nil {
		panic(err)
	}

	car1 := &domain.Car{
		Plate: "FMU526",
		Model: "Ford Fiesta",
	}
	err = db.Insert(car1)
	if err != nil {
		panic(err)
	}

	car2 := &domain.Car{
		Plate: "AAA111",
		Model: "Ford Focus",
	}
	err = db.Insert(car2)
	if err != nil {
		panic(err)
	}

	resident1 := &domain.Resident{
		Cars: []*domain.Car{car1},
		Name: "Matias Dominguez",
		Address: "Brandsen 123",
	}
	err = db.Insert(resident1)
	if err != nil {
		panic(err)
	}

	resident2 := &domain.Resident{
		Cars: []*domain.Car{car2},
		Name: "Pity Luz",
		Address: "Brandsen 123",
	}
	err = db.Insert(resident2)
	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&domain.Event{}, &domain.Resident{}, &domain.Car{}, &domain.User{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}