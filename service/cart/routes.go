package cart

import (
	"fmt"
	"net/http"

	"github.com/devdiagon/gomerce/service/auth"
	"github.com/devdiagon/gomerce/types"
	"github.com/devdiagon/gomerce/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	orderStore  types.OrderStore
	producStore types.ProductStore
	userStore   types.UserStore
}

func NewHandler(orderStore types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{orderStore: orderStore, producStore: productStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods("POST")
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	//Get user ID from the context
	userId := auth.GetUserIdFromContext(r.Context())

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

	//Business Logic implementation
	// get products Ids
	productIds, err := getCartItemsIds(cartPayload.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get the products based on their Ids
	products, err := h.producStore.GetProductsByIds(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// try to generate an Order
	orderId, totalPrice, err := h.createOrder(products, cartPayload.Items, userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// show the created Order if successful
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"totalPrice": totalPrice,
		"orderId":    orderId,
	})
}
