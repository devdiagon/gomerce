package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devdiagon/gomerce/types"
	"github.com/gorilla/mux"
)

type mockProductStore struct{}

func TestUserServiceHandlers(t *testing.T) {
	productStore := &mockProductStore{}
	handler := NewHandler(productStore)

	//TEST the product payload into de request
	t.Run("should fail if the product payload is invalid", func(t *testing.T) {
		//Create a dummy payload
		payload := types.CreateProductPayload{
			Name:        "test1",
			Description: "test1",
			Image:       "test1.jpg",
			Price:       3.40,
			Quantity:    5,
		}
		marshaled, _ := json.Marshal(payload)

		//Send the payload as a http request
		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshaled))
		if err != nil {
			t.Fatal(err)
		}

		//Send the request to de "/products" route to obtain a response
		res := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct)
		router.ServeHTTP(res, req)

		//Test the result
		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code: %d, got: %d", http.StatusBadRequest, res.Code)
		}
	})
}

// Mock data "from the handler"
func (m *mockProductStore) GetProducts() ([]types.Product, error) {
	return nil, fmt.Errorf("product not found")
}

func (m *mockProductStore) CreateProduct(product types.Product) error {
	return nil
}
