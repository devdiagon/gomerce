package cart

import (
	"net/http"

	"github.com/devdiagon/gomerce/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods("POST")
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {

}
