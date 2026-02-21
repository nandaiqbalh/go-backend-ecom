// Package types defines shared data structures and interfaces used across
// the service packages. This keeps types decoupled from specific
// implementations.
package types

// UserStore represents the minimum operations required by handlers and
// services to manage user records. Implementations may talk to a database,
// an in-memory store, or a remote service.
type UserStore interface {
    GetUserByEmail(email string) (*User, error)
    GetUserByID(id int) (*User, error)
    CreateUser(user *User) error 
}

// RegisterUserPayload defines the expected JSON structure for
// registration requests. Validation tags are used with the validator
// package to enforce required fields.
type RegisterUserPayload struct {
    FirstName string `json:"firstName" validate:"required"`
    LastName  string `json:"lastName" validate:"required"`
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6"`
}

// User represents a persisted user entity. The Password field is omitted
// from JSON serialization for security reasons.
type User struct {
    ID        int    `json:"id"`
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Email     string `json:"email"`
    Password  string `json:"-"`
    CreatedAt string `json:"createdAt"` 
} 