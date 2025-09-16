package api

import (
	"database/sql"
	"net/http"

	"github.com/devdiagon/gomerce/service/product"
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

	//Inject productStore into the service
	productStore := product.NewStore(s.db)
	productService := product.NewHandler(productStore)
	productService.RegisterRoutes(subrouter)

	log.Info("Server running on port ", s.address)

	//Send the port as ":<port>" string
	return http.ListenAndServe(s.address, router)
}
