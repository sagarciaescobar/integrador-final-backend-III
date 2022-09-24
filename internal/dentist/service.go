package dentist

import (
	"errors"
	"strconv"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
)

type Service interface {
	// GetByID busca un dentist por su id
	GetByID(id int) (domain.Dentist, error)
	// Create agrega un nuevo dentist
	Create(m map[string]string) (domain.Dentist, error)
	// Delete elimina un dentist
	DeleteById(id int) error
	// Update actualiza un dentist
	UpdateById(m map[string]string) (domain.Dentist, error)
	// Change Dentist Registration id
	ChangeRegistrationIdById(m map[string]string) error
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService() Service {
	return &service{r: NewRepository()}
}

func (s service) GetByID(id int) (domain.Dentist, error) {
	d, err := s.r.GetById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s service) Create(m map[string]string) (domain.Dentist, error) {
	d, err := Mapper(m)
	if err != nil {
		return domain.Dentist{}, err
	}
	d, err = s.r.Add(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s service) DeleteById(id int) error {
	err := s.r.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateById(m map[string]string) (domain.Dentist, error) {
	if m["id"] == "" {
		return domain.Dentist{}, errors.New("invalid id")
	}
	d, err := Mapper(m)
	if err != nil {
		return domain.Dentist{}, err
	}
	d, err = s.r.UpdateById(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s service) ChangeRegistrationIdById(m map[string]string) error {
	if m["id"] == "" {
		return errors.New("id must be defined")
	}
	if m["registration_id"] == "" {
		return errors.New("address must be defined")
	}
	id, err := strconv.Atoi(m["id"])
	if err != nil {
		return errors.New("invalid id")
	}
	err = s.r.ChangeRegistrationIdById(id, m["registration_id"])
	if err != nil {
		return err
	}
	return nil
}
