package domain

import (
	"encoding/json"
	"time"
)

// TODO IMPLEMENT

type User struct {
	Id int
	Username string
	Password string
	Role string
}

type Event struct {
	Id int
	Car *Car
	Resident *Resident
	CreationDate time.Time
}

type Resident struct {
	Id int
	Cars []*Car
	Name string
	Address string
}

func (resident Resident) toJson() (string, error) {
	ja, err := json.Marshal(resident)
	if err != nil {
		return "", err
	}
	return string(ja[:]), nil
}

type Car struct {
	Id int
	Plate string
	Model string
}

func (car Car) toJson() (string, error) {
	ja, err := json.Marshal(car)
	if err != nil {
		return "", err
	}
	return string(ja[:]), nil
}