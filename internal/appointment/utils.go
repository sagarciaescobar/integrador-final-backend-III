package appointment

import (
	"errors"
	"strconv"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/utils"
)

func Mapper(m map[string]string) (domain.Appointment, error) {
	if m["id_patient"] == "" {
		return domain.Appointment{}, errors.New("Appointment error - id_patient must be defined")
	} else if m["id_dentist"] == "" {
		return domain.Appointment{}, errors.New("Appointment error - id_dentist must be defined")
	} else if m["time"] == "" {
		return domain.Appointment{}, errors.New("Appointment error - time must be defined")
	}
	var a domain.Appointment
	if m["id"] != "" {
		id, err := strconv.Atoi(m["id"])
		if err != nil {
			return domain.Appointment{}, err
		}
		a.Id = id
	}
	idPatient, err := strconv.Atoi(m["id_patient"])
	if err != nil {
		return domain.Appointment{}, err
	}
	idDentist, err := strconv.Atoi(m["id_patient"])
	if err != nil {
		return domain.Appointment{}, err
	}
	a.IdPatient = idPatient
	a.IdDentist = idDentist

	data, err := utils.StringToDate(m["time"])
	if err != nil {
		return domain.Appointment{}, err
	}
	a.Time = data
	return a, nil
}
