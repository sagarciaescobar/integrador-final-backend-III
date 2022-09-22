package domain

import "time"

type Patient struct {
	Name             string    `json:"name"`
	LastName         string    `json:"lastName"`
	Address          string    `json:"address"`
	Dni              string    `json:"dni"`
	RegistrationDate time.Time `json:"registrationDate"`
}
