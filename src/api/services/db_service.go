package services

import (
	"github.com/matiasdominguez/pass/src/api/domain"
	"sync"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var dbInstance *pg.DB
var dbOnce sync.Once

func GetDBInstance() *pg.DB {
	dbOnce.Do(func() {
		dbInstance = pg.Connect(&pg.Options{
			User: "mdominguez",
			Database: "pass",
		})
	})
	return dbInstance
}


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
		Name: "Matias Dominguez",
		Address: "Brandsen 123",
	}
	err = db.Insert(resident1)
	if err != nil {
		panic(err)
	}

	resident2 := &domain.Resident{
		Name: "Pity Luz",
		Address: "Brandsen 123",
	}
	err = db.Insert(resident2)
	if err != nil {
		panic(err)
	}

	car1.ResidentId = resident1.Id
	err = db.Update(car1)
	if err != nil {
		panic(err)
	}

	car2.ResidentId = resident2.Id
	err = db.Update(car2)
	if err != nil {
		panic(err)
	}
}


func createSchema(db *pg.DB) error {
	orm.DropTable(db, &domain.Car{}, nil)
	orm.DropTable(db, &domain.Resident{}, nil)
	orm.DropTable(db, &domain.User{}, nil)
	orm.DropTable(db, &domain.Event{}, nil)
	for _, model := range []interface{}{&domain.Event{}, &domain.Resident{}, &domain.Car{}, &domain.User{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{IfNotExists:true})
		if err != nil {
			return err
		}
	}
	return nil
}