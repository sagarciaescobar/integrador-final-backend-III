package patient

import (
	"errors"
	"log"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/db"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/utils"
)

type Repository interface {
	// Add Patient
	Add(Patient domain.Patient) (domain.Patient, error)
	// Get Patient by Id
	GetById(id int) (domain.Patient, error)
	// Update Patient by Id
	UpdateById(domain.Patient) (domain.Patient, error)
	// Change Patient Address
	ChangeAddresById(id int, address string) error
	// Delete Patient
	DeleteById(id int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) Add(p domain.Patient) (domain.Patient, error) {
	Db := db.ConnDb()
	insertDb, err := Db.Prepare("INSERT INTO patient (name,last_name,address,dni,registration_date)VALUES (?,?,?,?,?)")
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
	if selDb.Next() {
		var id int
		var name, lastName, dni, address, date string
		selDb.Scan(&id, &name, &lastName, &address, &dni, &date)
		p.Id = id
		p.Name = name
		p.LastName = lastName
		p.Address = address
		p.Dni = dni
		data, err := utils.StringToDate(date)
		if err != nil {
			return domain.Patient{}, err
		}
		p.RegistrationDate = data
	} else {
		return domain.Patient{}, errors.New("patient not found")
	}
	defer Db.Close()
	log.Printf("[GET:patient] - name: %s - id: %d ", p.Name, p.Id)
	return p, nil
}

func (r repository) UpdateById(p domain.Patient) (domain.Patient, error) {
	Db := db.ConnDb()
	log.Println(p)
	selDb, err := Db.Prepare("UPDATE patient SET name=?,last_name=?,address=?,dni=?,registration_date=? WHERE id=?")
	if err != nil {
		return domain.Patient{}, err
	}
	Stmt, _ := selDb.Exec(p.Name, p.LastName, p.Address, p.Dni, p.RegistrationDate, p.Id)
	cant, err := Stmt.RowsAffected()
	if err != nil || cant != 1 {
		return domain.Patient{}, err
	}
	defer Db.Close()
	log.Printf("[Update:patient register] - cant: %d ", cant)
	return p, nil
}

func (r repository) ChangeAddresById(id int, address string) error {
	Db := db.ConnDb()
	selDb, err := Db.Prepare("UPDATE patient SET address=? WHERE id=?")
	if err != nil {
		return err
	}
	Stmt, _ := selDb.Exec(address, id)
	cant, err := Stmt.RowsAffected()
	if err != nil || cant != 1 {
		return err
	}
	defer Db.Close()
	log.Printf("[Update:patient address] - cant: %d ", cant)
	return nil
}

func (r repository) DeleteById(id int) error {
	Db := db.ConnDb()
	_, err := Db.Query("DELETE FROM patient WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
