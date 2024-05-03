package config

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	server   = "serve"
	user     = "user_testSQL"
	password = "Desarrollo1"
	dbname   = "Test_mgm"
	port     = "1433"
)

func DatabaseConnection() (db *sql.DB, err error) {
	sqlInfo := fmt.Sprintf("server=%sserver;user id=%s; password=%s,database=%s", server, user, password, dbname)
	db, err = sql.Open("mssql", sqlInfo)
	return
}
