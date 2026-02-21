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

type ProductStore interface {
	ListProducts() ([]*Product, error)
	GetProductByID(id int) (*Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id int) error
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

type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Image       string  `json:"image,omitempty"` // URL or path to product image; may be empty
    Price       float64 `json:"price"`
    Quantity    int     `json:"quantity"`
    CreatedAt   string  `json:"createdAt"`
}
