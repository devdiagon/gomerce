package api

import (
	"database/sql"
	"net/http"

	"github.com/devdiagon/gomerce/service/user"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		address: addr,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//Inject userStore into the service
	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(subrouter)

	log.Info("Server running on: ", s.address)

	return http.ListenAndServe(s.address, router)
}
