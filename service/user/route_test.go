package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nandaiqbalh/go-backend-ecom/types"
)

// TestUserServiceHandlers exercises the registration logic in
// the user handler. It uses a simple mock store that implements the
// UserStore interface so that database interactions are simulated.
func TestUserServiceHandlers(t *testing.T) {

    // mock store returns "not found" for any email
    userStore := &mockUserStore{}

    handler := NewHandler(userStore)

    t.Run("Should fail when request body is invalid", func(t *testing.T) {
        payload := types.RegisterUserPayload{
            FirstName: "John",
            LastName:  "Doe",
            Email:     "123", // invalid email according to validation tags
            Password:  "password123",
        }

        marshalled, _ := json.Marshal(payload)
        req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()

        router := mux.NewRouter()
        // register handler directly for simplicity
        router.HandleFunc("/register", handler.handleRegister)
        router.ServeHTTP(rr, req)

        if rr.Code != http.StatusBadRequest {
            t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
        }
    })

    t.Run("should correctly register user", func(t *testing.T) {
        payload := types.RegisterUserPayload{
            FirstName: "John",
            LastName:  "Doe",
            Email:     "valid@gmail.com",
            Password:  "password123",
        }

        marshalled, _ := json.Marshal(payload)
        req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()

        router := mux.NewRouter()
        router.HandleFunc("/register", handler.handleRegister)
        router.ServeHTTP(rr, req)

        if rr.Code != http.StatusCreated {
            t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
        }
    })
}

// mockUserStore satisfies types.UserStore with simple stubs.
// Each method returns fixed values to exercise different code paths.
type mockUserStore struct {}

func (m mockUserStore) GetUserByEmail(email string) (*types.User, error) {
    // always simulate user not found
    return nil, fmt.Errorf("user not found")
}

func (m mockUserStore) GetUserByID(id int) (*types.User, error) {
    return nil, nil
}

func (m mockUserStore) CreateUser(user *types.User) error {
    return nil
} 