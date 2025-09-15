package main

import (
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

	server := api.NewAPIServer(":8000", nil)
	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}
