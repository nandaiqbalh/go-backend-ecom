// Package api defines the HTTP server that exposes the application API.
// It wires together the database connection and request handlers, and sets
// up the routing structure using gorilla/mux.
package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandaiqbalh/go-backend-ecom/service/user"
)

// APIServer holds the address to listen on and a reference to the database
// connection used by handlers.
type APIServer struct {
    addr string
    db   *sql.DB
}

// NewAPIServer constructs a new APIServer with the given listen address and
// database connection. This allows the `main` package to configure the
// server and then start it.
func NewAPIServer(addr string, db *sql.DB) *APIServer {
    return &APIServer{
        addr: addr,
        db:   db,
    }
}

// Run starts the HTTP server. The flow is:
// 1. Create a root mux router.
// 2. Create a subrouter for versioned API paths (/api/v1).
// 3. Construct any required stores or services (here, the user store) and
//    inject them into handler constructors.
// 4. Register user-related routes on the subrouter.
// 5. Log the listening address and call http.ListenAndServe.
func (s *APIServer) Run() error {
    router := mux.NewRouter()

    // Use a versioned prefix to allow for future changes.
    subroute := router.PathPrefix("/api/v1").Subrouter()

    // Build a store backed by the shared database connection and pass it to
    // the user handler. This demonstrates dependency injection: the handler
    // need not know about the database itself, only the interface it needs.
    userStore := user.NewStore(s.db)
    userHandler := user.NewHandler(userStore)
    userHandler.RegisterRoutes(subroute)

    log.Println("Listening on", s.addr)

    return http.ListenAndServe(s.addr, router)
} 