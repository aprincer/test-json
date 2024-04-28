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
	sqlInfo := fmt.Sprintf("server=%server;user id=%s; password=%s", server, user, password)
	db, err = sql.Open("mssql", sqlInfo)
	return
}
