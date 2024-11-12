package models

import (
	"errors"
)

var ErrNorecord = errors.New("models: matche note not found")

type PersonData struct {
	ID          int
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"` // Я как понял это отчество
	Street      string `json:"street"`
	House       string `json:"house"`
	Building    string `json:"building"` // корпус
	Apartment   string `json:"apartment"`
	PhoneNumber string `json:"phone_number"`
}

type TestPerson struct {
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
}
