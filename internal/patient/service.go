package patient

import "github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"

type Service interface {
	// GetByID busca un dentist por su id
	GetByID(id int) (domain.Patient, error)
	// Create agrega un nuevo dentist
	//Create(p domain.Dentist) (domain.Dentist, error)
	// Delete elimina un producDentist
	//Delete(id int) error
	// Update actualiza un dentist
	//Update(id int, p domain.Dentist) (domain.Dentist, error)
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
