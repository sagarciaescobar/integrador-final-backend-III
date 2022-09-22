package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	mysql := Config.Application.Db.Mysql
	dataSource := mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/clinic"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Printf("connect with database [" + mysql.Port + "]")
	DB = StorageDB
	defer StorageDB.Close()
}
