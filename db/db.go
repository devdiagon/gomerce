package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
