package domain

import "time"

type Appointment struct {
	Id          int       `json:"id"`
	Patient     Patient   `json:"patient"`
	Dentist     Dentist   `json:"dentist"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}
