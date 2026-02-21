// Package user contains handlers and routing for user-related endpoints.
// This file defines a generic Handler struct which can later be extended
// with dependencies like a database connection or service layer.
package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandaiqbalh/go-backend-ecom/config"
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


// handleLogin processes login requests. The flow is:
// 1. Decode JSON body into LoginUserPayload.
// 2. Validate required fields (email + password).
// 3. Look up the user by email using the injected store.
// 4. Compare provided password with the stored hash.
// 5. Create a JWT containing the user ID.
// 6. Return the token in the response body.
//
// The handler deliberately returns a generic "invalid email or password"
// error for authentication failures to avoid giving attackers information
// about which part of the credentials was wrong.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
    if r.Body == nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("request body is empty"))
        return
    }

    var payload types.LoginUserPayload
    if err := utils.ParseJson(r, &payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    // validate payload using struct tags
    if err := utils.Validate.Struct(payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    // find user record
    u, err := h.store.GetUserByEmail(payload.Email)
    if err != nil {
        // do not reveal whether the email was wrong
        utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
        return
    }

    // check password
    if !auth.ComparePassword(u.Password, []byte(payload.Password)) {
        utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
        return
    }

    // ensure we have a signing secret
    if config.Envs.JWTSecret == "" {
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("JWT secret not configured"))
        return
    }
    secret := []byte(config.Envs.JWTSecret)

    token, err := auth.CreateJWT(secret, u.ID)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to generate token: %v", err))
        return
    }

    utils.WriteJson(w, http.StatusOK, types.LoginResponse{Token: token})
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