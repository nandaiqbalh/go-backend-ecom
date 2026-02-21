// Package user contains handlers and routing for user-related endpoints.
// This file defines a generic Handler struct which can later be extended
// with dependencies like a database connection or service layer.
package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandaiqbalh/go-backend-ecom/service/auth"
	"github.com/nandaiqbalh/go-backend-ecom/types"
	"github.com/nandaiqbalh/go-backend-ecom/utils"
)

// Handler is the HTTP handler for user operations. Dependencies like a
// database client or authentication service would be fields here.
type Handler struct {
    // You can add dependencies here, such as a database connection
	store types.UserStore
}

// NewHandler constructs a Handler. Dependencies can be initialized here.
func NewHandler(store types.UserStore) *Handler {
    return &Handler{
        store: store,
	}
}

// RegisterRoutes attaches user-related routes to the provided router. The
// API server calls this during startup to wire endpoints under the versioned
// subrouter.
func (h *Handler) RegisterRoutes(router *mux.Router) {
    // Define your user-related routes here
    router.HandleFunc("/login", h.handleLogin).Methods("POST")
    router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handleLogin processes login requests. In a real implementation, this
// would parse credentials, verify them against the database, and return a
// token or error.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
    // Implement login logic here
}

// handleRegister handles new user registration. Typical flow includes
// validating input, checking for existing users, hashing the password,
// creating the record, and sending back the created user or an error.
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
    // parse JSON payload from request body
    if r.Body == nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("request body is empty"))
        return
    }
    var payload types.RegisterUserPayload

    err := utils.ParseJson(r, &payload)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    // perform struct validation using validator tags defined in types
    if err := utils.Validate.Struct(payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    // verify user does not already exist
    _, err = h.store.GetUserByEmail(payload.Email)
    if err == nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
        return
    }

    // hash the password before saving
    hashPassword, err := auth.HashPassword(payload.Password)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to hash password: %v", err))
        return
    }

    // construct user object and persist it using the store
    user := &types.User{
        FirstName: payload.FirstName,
        LastName:  payload.LastName,
        Email:     payload.Email,
        Password:  hashPassword, // already hashed above
    }

    err = h.store.CreateUser(user)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    // respond with the newly created user data
    utils.WriteJson(w, http.StatusCreated, user)
} 