package patient

import (
	"errors"
	"strconv"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
)

type Service interface {
	// GetByID busca un dentist por su id
	GetByID(id int) (domain.Patient, error)
	// Create agrega un nuevo patient
	Create(m map[string]string) (domain.Patient, error)
	// Delete elimina un patient
	DeleteById(id int) error
	// Update actualiza un patient
	UpdateById(m map[string]string) (domain.Patient, error)
	// Change Patient Address
	ChangeAddresById(m map[string]string) error
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService() Service {
	return &service{r: NewRepository()}
}

func (s service) GetByID(id int) (domain.Patient, error) {
	p, err := s.r.GetById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s service) Create(m map[string]string) (domain.Patient, error) {
	p, err := Mapper(m)
	if err != nil {
		return domain.Patient{}, err
	}
	p, err = s.r.Add(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s service) DeleteById(id int) error {
	err := s.r.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateById(m map[string]string) (domain.Patient, error) {
	if m["id"] == "" {
		return domain.Patient{}, errors.New("id must be defined")
	}
	p, err := Mapper(m)
	if err != nil {
		return domain.Patient{}, err
	}
	p, err = s.r.UpdateById(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s service) ChangeAddresById(m map[string]string) error {
	if m["id"] == "" {
		return errors.New("id must be defined")
	}
	if m["address"] == "" {
		return errors.New("address must be defined")
	}
	id, err := strconv.Atoi(m["id"])
	if err != nil {
		return errors.New("invalid id")
	}
	err = s.r.ChangeAddresById(id, m["address"])
	if err != nil {
		return err
	}
	return nil
}
