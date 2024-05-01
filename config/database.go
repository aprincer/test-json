package config

import (
	"database/sql"
	"fmt"
)

const (
	server   = "localhost"
	user     = "sa"
	password = "password"
	dbname   = "database"
	port     = "port"
)

func DatabaseConnection() (db *sql.DB, err error) {
	sqlInfo := fmt.Sprintf("server=%sserver;user id=%s; password=%s,database=%s", server, user, password, dbname)
	db, err = sql.Open("mssql", sqlInfo)
	return
}
