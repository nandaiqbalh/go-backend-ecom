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