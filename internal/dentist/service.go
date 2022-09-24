package dentist

import (
	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
)

type Service interface {
	// GetByID busca un dentist por su id
	GetByID(id int) (domain.Dentist, error)
	// Create agrega un nuevo dentist
	Create(p domain.Dentist) (domain.Dentist, error)
	// Delete elimina un dentist
	DeleteById(id int) error
	// Update actualiza un dentist
	UpdateById(p domain.Dentist) (domain.Dentist, error)
	// Change Dentist Registration id
	ChangeRegistrationIdById(id int, rId string) error
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

func (s service) Create(d domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.Add(d)
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

func (s service) UpdateById(d domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.UpdateById(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s service) ChangeRegistrationIdById(id int, rId string) error {
	err := s.r.ChangeRegistrationIdById(id, rId)
	if err != nil {
		return err
	}
	return nil
}
