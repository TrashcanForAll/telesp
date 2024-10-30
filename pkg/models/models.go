package models

import (
	"errors"
)

var ErrNorecord = errors.New("models: matche note not found")

type PersonData struct {
	ID int
	LastName string
	FirstName string
	MiddleName string // Я как понял это отчество
	Street string
	House int
	Housing int // корпус
	Flat int
	PhoneNumber int
}