package config

import (
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysql := Config.Application.Db.Mysql
	dataSource := mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/clinic"
}
