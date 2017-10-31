package domain

import (
	"encoding/json"
	"time"
	"fmt"
)

type User struct {
	Id       int
	Username string
	Password string
	Role     string
}

type Event struct {
	Id           int
	Car          *Car
	Resident     *Resident
	CreationDate time.Time
}

type Resident struct {
	Id      int
	Cars    []Car `json:"car"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (resident Resident) toJson() (string, error) {
	ja, err := json.Marshal(resident)
	if err != nil {
		return "", err
	}
	return string(ja[:]), nil
}

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
}

type Car struct {
	Plate     string `sql:",pk" json:"plate"`
	Model     string `json:"model"`
	Residents []Resident `json:"resident"`
}

func (car Car) toJson() (string, error) {
	ja, err := json.Marshal(car)
	if err != nil {
		return "", err
	}
	return string(ja[:]), nil
}
