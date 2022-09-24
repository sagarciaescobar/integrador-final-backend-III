package dentist

import (
	"errors"
	"strconv"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
)

func Mapper(m map[string]string) (domain.Dentist, error) {
	if m["name"] == "" {
		return domain.Dentist{}, errors.New("Dentist error - name must be defined")
	} else if m["last_name"] == "" {
		return domain.Dentist{}, errors.New("Dentist error - last_name must be defined")
	} else if m["registration_id"] == "" {
		return domain.Dentist{}, errors.New("Dentist error - registration_id must be defined")
	}
	var d domain.Dentist
	if m["id"] != "" {
		id, err := strconv.Atoi(m["id"])
		if err != nil {
			return domain.Dentist{}, err
		}
		d.Id = id
	}
	d.Name = m["name"]
	d.LastName = m["last_name"]
	d.RegistrationId = m["registration_id"]

	return d, nil
}
