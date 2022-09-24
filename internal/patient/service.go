package patient

import (
	"log"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
)

type Service interface {
	// GetByID busca un dentist por su id
	GetByID(id int) (domain.Patient, error)
	// Create agrega un nuevo patient
	Create(p domain.Patient) (domain.Patient, error)
	// Delete elimina un patient
	DeleteById(id int) error
	// Update actualiza un patient
	UpdateById(p domain.Patient) (domain.Patient, error)
	// Change Patient Address
	ChangeAddresById(id int, address string) error
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

func (s service) Create(p domain.Patient) (domain.Patient, error) {
	p, err := s.r.Add(p)
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

func (s service) UpdateById(p domain.Patient) (domain.Patient, error) {
	log.Println(p)
	p, err := s.r.UpdateById(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s service) ChangeAddresById(id int, address string) error {
	err := s.r.ChangeAddresById(id, address)
	if err != nil {
		return err
	}
	return nil
}
