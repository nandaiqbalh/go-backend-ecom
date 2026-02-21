package types

// LoginResponse represents the JSON body returned to a successful login.
// Only the token is sent; user details may be added if needed later.
type LoginResponse struct {
    Token string `json:"token"`
}