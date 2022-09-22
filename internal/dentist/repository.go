package dentist

import (
	"database/sql"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
)

type Repository interface {
	Add(Patient domain.Patient) domain.Patient
	GetById(id int) domain.Patient
	UpdateById(domain.Patient) domain.Patient
	ChangeNameById(domain.Patient) domain.Patient
	Delete(id int) error
}

type db struct {
	store *sql.DB
}

func (*db) NewRepository() *sql.DB {
	return &sql.DB{}
}

func (*db) Add(Patient domain.Patient) domain.Patient {
}

func (*db) 
