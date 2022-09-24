package domain

import "time"

type Appointment struct {
	Id          int       `json:"id"`
	IdPatient   int       `json:"id_patient"`
	IdDentist   int       `json:"id_dentist"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}
