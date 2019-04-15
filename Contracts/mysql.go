package contracts

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

//Connect sqldb
func Connect() *sql.DB {
	dbConfig := mysql.NewConfig()
	dbConfig.User = "docker"
	dbConfig.Passwd = "docker"
	dbConfig.Addr = "mysql:3306"
	dbConfig.DBName = "kumparannews"
	dbConfig.Net = "tcp"
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}
