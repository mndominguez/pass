package domain

import (
	"encoding/json"
	"time"
	//"fmt"
)

type User struct {
	Id       int
	Username string
	Password string
	Role     string
}

type Event struct {
	Id           int
	CarID        int `pg:",fk:Car"`
	ResidentID   int `pg:",fk:Resident"`
	CreationDate time.Time
}

type Resident struct {
	Id      int
	Name    string `json:"name"`
	Address string `json:"address"`
	Cars []*Car `json:"cars"`
}

type Car struct {
	Id 		  int
	Plate     string `json:"plate"`
	Model     string `json:"model"`
	ResidentId int `json:"resident"`
}

func (resident Resident) toJson() (string, error) {
	ja, err := json.Marshal(resident)
	if err != nil {
		return "", err
	}
	return string(ja[:]), nil
}
/*
func (c *Car) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&c.Plate, &c.Model}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
	}
	return nil
}*/

func (car Car) toJson() (string, error) {
	ja, err := json.Marshal(car)
	if err != nil {
		return "", err
	}
	return string(ja[:]), nil
}
