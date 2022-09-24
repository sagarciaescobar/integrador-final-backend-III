package appointment

import (
	"errors"
	"strconv"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/utils"
)

type Service interface {
	// GetByID busca un dentist por su id
	GetByID(id int) (domain.Appointment, error)
	// Create agrega un nuevo appointment
	Create(m map[string]string) (domain.Appointment, error)
	// Delete elimina un appointment
	DeleteById(id int) error
	// Update actualiza un appointment
	UpdateById(m map[string]string) (domain.Appointment, error)
	// Change Appointment time
	ChangeTimeById(m map[string]string) error
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService() Service {
	return &service{r: NewRepository()}
}

func (s service) GetByID(id int) (domain.Appointment, error) {
	p, err := s.r.GetById(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return p, nil
}

func (s service) Create(m map[string]string) (domain.Appointment, error) {
	data, err := Mapper(m)
	if err != nil {
		return domain.Appointment{}, err
	}
	a, err := s.r.Add(data)
	if err != nil {
		return domain.Appointment{}, err
	}
	return a, nil
}

func (s service) DeleteById(id int) error {
	err := s.r.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateById(m map[string]string) (domain.Appointment, error) {
	if m["id"] == "" {
		return domain.Appointment{}, errors.New("must be id in petition")
	}
	data, err := Mapper(m)
	if err != nil {
		return domain.Appointment{}, err
	}
	a, err := s.r.UpdateById(data)
	if err != nil {
		return domain.Appointment{}, err
	}
	return a, nil
}

func (s service) ChangeTimeById(m map[string]string) error {
	if m["id"] == "" || m["time"] == "" {
		return errors.New("id and time must be defined")
	}
	id, err := strconv.Atoi(m["id"])
	if err != nil {
		return errors.New("invalid id")
	}
	time, err := utils.StringToDate(m["time"])
	if err != nil {
		return err
	}
	err = s.r.ChangeTimeById(id, time)
	if err != nil {
		return err
	}
	return nil
}
