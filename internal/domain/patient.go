package domain

import "time"

type Patient struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	LastName         string    `json:"lastName"`
	Address          string    `json:"address"`
	Dni              string    `json:"dni"`
	RegistrationDate time.Time `json:"registrationDate"`
}
