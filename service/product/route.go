package product

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/nandaiqbalh/go-backend-ecom/service/auth"
    "github.com/nandaiqbalh/go-backend-ecom/types"
    "github.com/nandaiqbalh/go-backend-ecom/utils"
)

type Handler struct {
    store types.ProductStore
}

// NewHandler creates a new Handler with the given ProductStore.
func NewHandler(store types.ProductStore) *Handler {
    return &Handler{store: store}
}

// RegisterRoutes attaches product-related routes to the provided router.
// All product endpoints require a valid JWT bearer token; we apply the
// authentication middleware here.
func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/products", auth.RequireToken(h.handleListProducts)).Methods("GET")
    router.HandleFunc("/products", auth.RequireToken(h.handleCreateProduct)).Methods("POST")
    router.HandleFunc("/products/{id}", auth.RequireToken(h.handleGetProduct)).Methods("GET")
    router.HandleFunc("/products/{id}", auth.RequireToken(h.handleUpdateProduct)).Methods("PUT")
    router.HandleFunc("/products/{id}", auth.RequireToken(h.handleDeleteProduct)).Methods("DELETE")
}

func (h *Handler) handleListProducts(w http.ResponseWriter, r *http.Request) {
    products, err := h.store.ListProducts()
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to list products: %v", err))
        return
    }

    utils.WriteJson(w, http.StatusOK, products)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
    if r.Body == nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("request body is empty"))
        return
    }

    var payload types.CreateProductPayload
    if err := utils.ParseJson(r, &payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    if err := utils.Validate.Struct(payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    prod := &types.Product{
        Name:        payload.Name,
        Description: payload.Description,
        Price:       payload.Price,
        Quantity:    payload.Quantity,
    }

    if err := h.store.CreateProduct(prod); err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    resp := types.CreateProductResponse{
        Message: "product created",
        Data: struct {
            ID          int     `json:"id"`
            Name        string  `json:"name"`
            Description string  `json:"description"`
            Price       float64 `json:"price"`
            Quantity    int     `json:"quantity"`
            CreatedAt   string  `json:"createdAt"`
        }{
            ID:          prod.ID,
            Name:        prod.Name,
            Description: prod.Description,
            Price:       prod.Price,
            Quantity:    prod.Quantity,
            CreatedAt:   prod.CreatedAt,
        },
    }
    utils.WriteJson(w, http.StatusCreated, resp)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
        return
    }

    prod, err := h.store.GetProductByID(id)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }
    if prod == nil {
        utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
        return
    }

    resp := types.GetProductByIDResponse{
        Message: "success",
        Data: struct {
            ID          int     `json:"id"`
            Name        string  `json:"name"`
            Description string  `json:"description"`
            Price       float64 `json:"price"`
            Quantity    int     `json:"quantity"`
            CreatedAt   string  `json:"createdAt"`
        }{
            ID:          prod.ID,
            Name:        prod.Name,
            Description: prod.Description,
            Price:       prod.Price,
            Quantity:    prod.Quantity,
            CreatedAt:   prod.CreatedAt,
        },
    }
    utils.WriteJson(w, http.StatusOK, resp)
}

func (h *Handler) handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
        return
    }

    if r.Body == nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("request body is empty"))
        return
    }

    var payload types.UpdateProductPayload
    if err := utils.ParseJson(r, &payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    if err := utils.Validate.Struct(payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    if payload.ID != id {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id mismatch"))
        return
    }

    prod := &types.Product{
        ID:          id,
        Name:        payload.Name,
        Description: payload.Description,
        Price:       payload.Price,
        Quantity:    payload.Quantity,
    }

    if err := h.store.UpdateProduct(prod); err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    resp := types.UpdateProductResponse{
        Message: "product updated",
        Data: struct {
            ID          int     `json:"id"`
            Name        string  `json:"name"`
            Description string  `json:"description"`
            Price       float64 `json:"price"`
            Quantity    int     `json:"quantity"`
            CreatedAt   string  `json:"createdAt"`
        }{
            ID:          prod.ID,
            Name:        prod.Name,
            Description: prod.Description,
            Price:       prod.Price,
            Quantity:    prod.Quantity,
            CreatedAt:   prod.CreatedAt,
        },
    }
    utils.WriteJson(w, http.StatusOK, resp)
}

func (h *Handler) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
        return
    }

    if err := h.store.DeleteProduct(id); err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    resp := types.DeleteProductResponse{
        Message: "product deleted",
        Data: struct {
            ID          int     `json:"id"`
            Name        string  `json:"name"`
            Description string  `json:"description"`
            Price       float64 `json:"price"`
            Quantity    int     `json:"quantity"`
            CreatedAt   string  `json:"createdAt"`
        }{},
    }
    utils.WriteJson(w, http.StatusOK, resp)
}
