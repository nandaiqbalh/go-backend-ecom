package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nandaiqbalh/go-backend-ecom/config"
	"github.com/nandaiqbalh/go-backend-ecom/utils"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	// Implement JWT creation logic here

	expiration := time.Second *time.Duration(config.Envs.JWTExpirationSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.Itoa(userId),
		"exp":    time.Now().Add(expiration).Unix(), // token expires in 72 hours
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
// ParseJWT verifies the provided token string using the secret and returns
// the claims if the token is valid. It returns an error otherwise.
func ParseJWT(tokenString string, secret []byte) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        // ensure the signing method is HMAC so we don't accept other types
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return secret, nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, fmt.Errorf("invalid token")
}

// RequireToken is an HTTP middleware that checks for a Bearer token in the
// Authorization header. If the token is valid the request is forwarded to
// the next handler and the user ID (if present) is stored in the context.
func RequireToken(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("missing Authorization header"))
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid Authorization header"))
            return
        }

        secret := []byte(config.Envs.JWTSecret)
        claims, err := ParseJWT(parts[1], secret)
        if err != nil {
            utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: %v", err))
            return
        }

        // attach userId to context if present
        if uid, ok := claims["userId"].(string); ok {
            if id, err := strconv.Atoi(uid); err == nil {
                ctx := context.WithValue(r.Context(), "userId", id)
                r = r.WithContext(ctx)
            }
        }

        next(w, r)
    }
}
