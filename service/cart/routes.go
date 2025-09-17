package cart

import (
	"fmt"
	"net/http"

	"github.com/devdiagon/gomerce/types"
	"github.com/devdiagon/gomerce/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	orderStore  types.OrderStore
	producStore types.ProductStore
}

func NewHandler(orderStore types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{orderStore: orderStore, producStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods("POST")
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	//Get data from the Body as JSON
	var cartPayload types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cartPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Validate the data sent
	if err := utils.Validate.Struct(cartPayload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

}
