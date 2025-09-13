package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
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

	return http.ListenAndServe(s.address, router)
}
