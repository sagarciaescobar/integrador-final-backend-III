package patient

import (
	"errors"
	"strconv"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/utils"
)

func Mapper(m map[string]string) (domain.Patient, error) {
	if m["name"] == "" {
		return domain.Patient{}, errors.New("Patient error - name must be defined")
	} else if m["last_name"] == "" {
		return domain.Patient{}, errors.New("Patient error - last_name must be defined")
	} else if m["address"] == "" {
		return domain.Patient{}, errors.New("Patient error - address must be defined")
	} else if m["dni"] == "" {
		return domain.Patient{}, errors.New("Patient error - dni must be defined")
	} else if m["registration_date"] == "" {
		return domain.Patient{}, errors.New("Patient error - registration_date must be defined")
	}
	var p domain.Patient
	if m["id"] != "" {
		id, err := strconv.Atoi(m["id"])
		if err != nil {
			return domain.Patient{}, err
		}
		p.Id = id
	}
	p.Name = m["name"]
	p.LastName = m["last_name"]
	p.Address = m["address"]
	p.Dni = m["dni"]

	data, err := utils.StringToDate(m["registration_date"])
	if err != nil {
		return domain.Patient{}, err
	}
	p.RegistrationDate = data
	return p, nil
}
