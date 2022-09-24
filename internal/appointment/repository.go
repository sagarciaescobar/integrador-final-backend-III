package appointment

import (
	"errors"
	"log"
	"time"

	"github.com/sagarciaescobar/integrador-final-backend-III/internal/domain"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/db"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/utils"
)

type Repository interface {
	// Add Appointment
	Add(a domain.Appointment) (domain.Appointment, error)
	// Get Appointment by Id
	GetById(id int) (domain.Appointment, error)
	// Update Appointment by Id
	UpdateById(domain.Appointment) (domain.Appointment, error)
	// Change Appointment Address
	ChangeTimeById(id int, time time.Time) error
	// Delete Appointment
	DeleteById(id int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) Add(a domain.Appointment) (domain.Appointment, error) {
	Db := db.ConnDb()
	insertDb, err := Db.Prepare("INSERT INTO appointment (id_patient,id_dentist,time,description)VALUES (?,?,?,?)")
	if err != nil {
		return domain.Appointment{}, err
	}
	Stmt, err := insertDb.Exec(a.IdPatient, a.IdDentist, a.Time, a.Description)
	if err != nil {
		return domain.Appointment{}, err
	}
	id, _ := Stmt.LastInsertId()
	a.Id = int(id)
	log.Printf("[ADD:appointment] - patient: %d - dentist: %d ", a.IdPatient, a.IdPatient)
	return a, nil
}

func (r repository) GetById(id int) (domain.Appointment, error) {
	Db := db.ConnDb()
	selDb, err := Db.Query("SELECT * FROM appointment WHERE id=?", id)
	if err != nil {
		return domain.Appointment{}, err
	}
	a := domain.Appointment{}
	if selDb.Next() {
		var id, idPatient, idDentist int
		var date, description string
		selDb.Scan(&id, &idPatient, &idDentist, &date, &description)
		a.Id = id
		a.IdPatient = idPatient
		a.IdDentist = idDentist
		data, err := utils.StringToDate(date)
		if err != nil {
			return domain.Appointment{}, err
		}
		a.Time = data
	} else {
		return domain.Appointment{}, errors.New("appointment not found")
	}
	defer Db.Close()
	log.Printf("[GET:appointment] - patient: %d - dentist: %d ", a.IdPatient, a.IdPatient)
	return a, nil
}

func (r repository) UpdateById(a domain.Appointment) (domain.Appointment, error) {
	Db := db.ConnDb()
	selDb, err := Db.Prepare("UPDATE appointment SET id_patient=?,id_dentist=?,time=?,description=? WHERE id=?")
	if err != nil {
		return domain.Appointment{}, err
	}
	Stmt, _ := selDb.Exec(a.IdPatient, a.IdDentist, a.Time, a.Description, a.Id)
	cant, err := Stmt.RowsAffected()
	if err != nil || cant != 1 {
		return domain.Appointment{}, err
	}
	defer Db.Close()
	log.Printf("[Update:appointment register] - cant: %d ", cant)
	return a, nil
}

func (r repository) ChangeTimeById(id int, time time.Time) error {
	Db := db.ConnDb()
	selDb, err := Db.Prepare("UPDATE patient SET time=? WHERE id=?")
	if err != nil {
		return err
	}
	Stmt, _ := selDb.Exec(time, id)
	cant, err := Stmt.RowsAffected()
	if err != nil || cant != 1 {
		return err
	}
	defer Db.Close()
	log.Printf("[Update:appointment time] - cant: %d ", cant)
	return nil
}

func (r repository) DeleteById(id int) error {
	Db := db.ConnDb()
	_, err := Db.Query("DELETE FROM appointment WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
