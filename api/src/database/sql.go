package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Init - Open connection bd sql and return
func Init() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnetion)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
