package domain

import "time"

type Appointment struct {
	Patient     Patient   `json:"patient"`
	Dentist     Dentist   `json:"dentist"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}
