package types

// RegisterUserPayload defines the expected JSON structure for
// registration requests. Validation tags are used with the validator
// package to enforce required fields.
type RegisterUserPayload struct {
    FirstName string `json:"firstName" validate:"required"`
    LastName  string `json:"lastName" validate:"required"`
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6"`
}

// LoginUserPayload represents the JSON structure for login requests. Both
// fields are required and validated by the `validator` instance in utils.
type LoginUserPayload struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,gte=0"`
}

type GetProductByIDPayload struct {
	ID int `json:"id" validate:"required"`
}

type UpdateProductPayload struct {
	ID          int     `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,gte=0"`
}

type DeleteProductPayload struct {
	ID int `json:"id" validate:"required"`
}
