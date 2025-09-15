package user

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

type mockUserStore struct{}

func TestUserServiceHandlers(t *testing.T) {
	//Create the handler from the routes that we're using, in this case they belong to "UserStore" type
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	//TEST the user payload into de request
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		//Create a dummy payload
		payload := types.RegisterUserPayload{
			FirstName: "test1",
			LastName:  "test1",
			Email:     "invalid",
			Password:  "ABC123",
		}
		marshaled, _ := json.Marshal(payload)

		//Send the payload as a http request
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshaled))
		if err != nil {
			t.Fatal(err)
		}

		//Send the request to de "/register" route to obtain a response
		res := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(res, req)

		//Test the result
		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code: %d, got: %d", http.StatusBadRequest, res.Code)
		}
	})

	t.Run("should register the user", func(t *testing.T) {
		//Create a dummy payload
		payload := types.RegisterUserPayload{
			FirstName: "test1",
			LastName:  "test1",
			Email:     "valid@mail.com",
			Password:  "ABC123",
		}
		marshaled, _ := json.Marshal(payload)

		//Send the payload as a http request
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshaled))
		if err != nil {
			t.Fatal(err)
		}

		//Send the request to de "/register" route to obtain a response
		res := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(res, req)

		//Test the result
		if res.Code != http.StatusCreated {
			t.Errorf("expected status code: %d, got: %d", http.StatusBadRequest, res.Code)
		}
	})
}

// Mock data "from the handler"
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
