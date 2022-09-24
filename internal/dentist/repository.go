package dentist

import (
	"errors"
	"log"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/db"
)

type Repository interface {
	// Add dentist
	Add(Patient domain.Dentist) (domain.Dentist, error)
	// Get dentist by Id
	GetById(id int) (domain.Dentist, error)
	// Update dentist by Id
	UpdateById(domain.Dentist) (domain.Dentist, error)
	// Change dentist Address
	ChangeRegistrationIdById(id int, r string) error
	// Delete dentist
	DeleteById(id int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) Add(d domain.Dentist) (domain.Dentist, error) {
	Db := db.ConnDb()
	insertDb, err := Db.Prepare("INSERT INTO dentist (name,last_name,registration_id)VALUES (?,?,?)")
	if err != nil {
		return domain.Dentist{}, err
	}
	Stmt, err := insertDb.Exec(d.Name, d.LastName, d.RegistrationId)
	if err != nil {
		return domain.Dentist{}, err
	}
	id, _ := Stmt.LastInsertId()
	d.Id = int(id)
	log.Printf("[ADD:dentist] - name: %s - id: %d ", d.Name, d.Id)
	return d, nil
}

func (r repository) GetById(id int) (domain.Dentist, error) {
	Db := db.ConnDb()
	selDb, err := Db.Query("SELECT * FROM dentist WHERE id=?", id)
	if err != nil {
		return domain.Dentist{}, err
	}
	d := domain.Dentist{}
	if selDb.Next() {
		var id int
		var name, lastName, registrationId string
		selDb.Scan(&id, &name, &lastName, &registrationId)
		d.Id = id
		d.Name = name
		d.LastName = lastName
		d.RegistrationId = registrationId
	} else {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	defer Db.Close()
	log.Printf("[GET:dentist] - name: %s - id: %d ", d.Name, d.Id)
	return d, nil
}

func (r repository) UpdateById(d domain.Dentist) (domain.Dentist, error) {
	Db := db.ConnDb()
	selDb, err := Db.Prepare("UPDATE dentist SET name=?,last_name=?,registration_id=? WHERE id=?")
	if err != nil {
		return domain.Dentist{}, err
	}
	Stmt, _ := selDb.Exec(d.Name, d.LastName, d.RegistrationId)
	cant, err := Stmt.RowsAffected()
	if err != nil || cant != 1 {
		return domain.Dentist{}, err
	}
	defer Db.Close()
	log.Printf("[Update:dentist register] - cant: %d ", cant)
	return d, nil
}

func (r repository) ChangeRegistrationIdById(id int, rId string) error {
	Db := db.ConnDb()
	selDb, err := Db.Prepare("UPDATE dentist SET registration_id=? WHERE id=?")
	if err != nil {
		return err
	}
	Stmt, _ := selDb.Exec(rId, id)
	cant, err := Stmt.RowsAffected()
	if err != nil || cant != 1 {
		return err
	}
	defer Db.Close()
	log.Printf("[Update:dentist registration id] - cant: %d ", cant)
	return nil
}

func (r repository) DeleteById(id int) error {
	Db := db.ConnDb()
	_, err := Db.Query("DELETE FROM dentist WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
