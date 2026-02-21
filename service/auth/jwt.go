package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nandaiqbalh/go-backend-ecom/config"
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