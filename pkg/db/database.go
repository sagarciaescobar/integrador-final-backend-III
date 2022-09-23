package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sagarciaescobar/integrador-final-backend-III/config"
)

func Test() {
	mysql := config.Config.Application.Db.Mysql
	dataSource := mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/clinic"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Printf("connect with database [" + mysql.Port + "]")
	defer StorageDB.Close()
}

func ConnDb() *sql.DB {
	mysql := config.Config.Application.Db.Mysql
	dataSource := mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/clinic"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
