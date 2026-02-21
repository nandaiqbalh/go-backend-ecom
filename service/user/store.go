// Package user provides data access logic for user records stored in the
// database. Store wraps an *sql.DB and exposes methods matching the
// UserStore interface defined in the `types` package.
package user

import (
	"database/sql"
	"fmt"

	"github.com/nandaiqbalh/go-backend-ecom/types"
)

// Store holds a SQL database connection.
type Store struct {
    db *sql.DB
}

// NewStore creates a Store given a previously opened *sql.DB.
func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

// GetUserByEmail queries the database for a user with the given email. If no
// rows are returned, it returns an error indicating the user was not found.
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
    rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
    if err != nil {
        return nil, err
    }

    u := new(types.User)

    for rows.Next() {
        u, err = ScanRowIntoUser(rows)
        if err != nil {
            return nil, err
        }
    }

    if u.ID == 0 {
        return nil, fmt.Errorf("user not found")
    }

    return u, nil
}

// ScanRowIntoUser is a helper that reads the current row from *sql.Rows into a
// types.User struct. It abstracts the scanning logic used by multiple
// methods.
func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
    user := new(types.User)

    err := rows.Scan(
        &user.ID,
        &user.FirstName,
        &user.LastName,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
    )

    if err != nil {
        return nil, err
    }

    return user, nil
}

// CreateUser inserts a new user record into the database. Caller must ensure
// the User struct has valid data; password should already be hashed.
func (s *Store) CreateUser(user *types.User) error {
    _, err := s.db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)",
        user.FirstName, user.LastName, user.Email, user.Password)

    return err
}

// GetUserByID returns a user matching the given ID or an error if none exists.
func (s *Store) GetUserByID(id int) (*types.User, error) {
    rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
    if err != nil {
        return nil, err
    }

    u := new(types.User)

    for rows.Next() {
        u, err = ScanRowIntoUser(rows)
        if err != nil {
            return nil, err
        }
    }

    if u.ID == 0 {
        return nil, fmt.Errorf("user not found")
    }

    return u, nil
} 