package main

import (
	"github.com/devdiagon/gomerce/cmd/api"
	"github.com/devdiagon/gomerce/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := db.NewMySQLStorage()

	server := api.NewAPIServer(":8000", nil)
	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}
