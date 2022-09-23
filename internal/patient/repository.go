package patient

import (
	"log"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/db"
)

type Repository interface {
	// Add Patient
	Add(Patient domain.Patient) (domain.Patient, error)
	// Get Patient by Id
	GetById(id int) (domain.Patient, error)
	// Update Patient by Id
	//UpdateById(domain.Patient) (domain.Patient, error)
	// Change Patient Address
	//ChangeAddresById(address string) (domain.Patient, error)
	// Delete Patient
	//Delete(id int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) Add(p domain.Patient) (domain.Patient, error) {
	Db := db.ConnDb()
	insertDb, err := Db.Prepare("INSERT INTO patient (?,?,?,?)")
	if err != nil {
		return domain.Patient{}, err
	}
	Stmt, err := insertDb.Exec(p.Name, p.LastName, p.Address, p.Dni, p.RegistrationDate)
	if err != nil {
		return domain.Patient{}, err
	}
	id, _ := Stmt.LastInsertId()
	p.Id = int(id)
	log.Printf("[ADD:patient] - name: %s - id: %d ", p.Name, p.Id)
	return p, nil
}

func (r repository) GetById(id int) (domain.Patient, error) {
	Db := db.ConnDb()
	selDb, err := Db.Query("SELECT * FROM patient WHERE id=?", id)
	if err != nil {
		return domain.Patient{}, err
	}
	p := domain.Patient{}
	for selDb.Next() {
		var id int
		var name, lastName, dni, address, date string
		selDb.Scan(&id, &name, &lastName, &address, &dni, &date)
		p.Id = id
		p.Name = name
		p.LastName = lastName
		p.Address = address
		p.Dni = dni
		log.Println(date)
	}
	log.Printf("[ADD:patient] - name: %s - id: %d ", p.Name, p.Id)
	return p, nil
}
