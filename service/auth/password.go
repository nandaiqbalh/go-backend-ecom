// Package auth contains authentication utility functions, such as password
// hashing.
package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plaintext password and returns a bcrypt hash. In a
// real application this would also consider peppering or cost configuration.
func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

    if err != nil {
        return "", err
    }

    return string(hash), nil
} 