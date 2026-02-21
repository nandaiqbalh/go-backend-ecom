package types

// LoginResponse represents the JSON body returned to a successful login.
// Only the token is sent; user details may be added if needed later.
type LoginResponse struct {
    Token string `json:"token"`
}

type CreateProductResponse struct {
	Message string `json:"message"`
	Data	struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Price float64 `json:"price"`
		Quantity int `json:"quantity"`
		CreatedAt string `json:"createdAt"`
	} `json:"data"`
}

type GetProductByIDResponse struct {
	Message string `json:"message"`
	Data	struct {
		ID int `json:"id"`
		Name string `json:"name"`	
		Description string `json:"description"`
		Price float64 `json:"price"`
		Quantity int `json:"quantity"`
		CreatedAt string `json:"createdAt"`
	} `json:"data"`
}

type UpdateProductResponse struct {
	Message string `json:"message"`
	Data	struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Price float64 `json:"price"`
		Quantity int `json:"quantity"`
		CreatedAt string `json:"createdAt"`
	} `json:"data"`
}

type DeleteProductResponse struct {
	Message string `json:"message"`
	Data	struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Price float64 `json:"price"`
		Quantity int `json:"quantity"`
		CreatedAt string `json:"createdAt"`	
	} `json:"data"`
}

