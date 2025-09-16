package product

import (
	"net/http"

	"github.com/devdiagon/gomerce/types"
	"github.com/devdiagon/gomerce/utils"
	"github.com/gorilla/mux"
)

// Dependency injection done with the "Handler"
type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods("GET")
	router.HandleFunc("/products", h.handleCreateProduct).Methods("POST")
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {}
