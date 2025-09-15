package main

import (
	"database/sql"

	"github.com/devdiagon/gomerce/cmd/api"
	"github.com/devdiagon/gomerce/config"
	"github.com/devdiagon/gomerce/db"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBADDRESS,
		DBName:               config.Envs.DBNAME,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Successfully connected to the Database")
}
