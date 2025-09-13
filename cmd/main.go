package main

import (
	"github.com/devdiagon/gomerce/cmd/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	server := api.NewAPIServer(":8000", nil)
	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}
