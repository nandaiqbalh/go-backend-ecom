// Package config handles application configuration values sourced from
// environment variables. The values are read at startup and exposed via the
// public `Envs` variable so other packages can access them.
//
// The configuration struct mirrors the expected environment variables for
// things like the public host, port, and database credentials.
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration values needed throughout the application.
// Fields have reasonable fallbacks in case the corresponding environment
// variables are not set, which makes local development easier.
//
// Typical flow:
// 1. The package initializes via the `initConfig` function when imported.
// 2. `godotenv.Load` is called to load values from a .env file (if present).
// 3. `getEnv` retrieves each environment variable with a fallback.
// 4. The resulting Config struct is stored in the package-level `Envs`.

// Config contains configuration values read from the environment.
type Config struct {
    PublicHost string
    Port       string

    DBUser     string
    DBPassword string
    DBAddress  string
    DBName     string

	JWTExpirationSeconds int64
	JWTSecret            string
}

// Envs is the globally accessible configuration populated during init.
var Envs = initConfig()

// initConfig reads environment variables (optionally from a .env file) and
// returns a Config struct with all values filled in. This function is
// executed once during package initialization.
func initConfig() Config {
    // Load variables defined in a .env file into the environment. Errors are
    // ignored because absence of a .env file is not fatal.
    godotenv.Load()

    return Config{
        PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
        Port:       getEnv("PORT", "8080"),
        DBUser:     getEnv("DB_USER", "root"),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
        DBName:     getEnv("DB_NAME", "go-backend-ecom"),
		JWTExpirationSeconds: getEnvAsInt("JWT_EXPIRATION_SECONDS", 3600*24), // default to 24 hours
		JWTSecret:            getEnv("JWT_SECRET", "asdfasdfasdf"), // default to a placeholder secret; should be overridden in production
    }
}

// getEnv returns the value of the environment variable named by the key.
// If the variable is not present, it returns the provided fallback.
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if valueStr, ok := os.LookupEnv(key); ok {
		var value int
		_, err := fmt.Sscanf(valueStr, "%d", &value)
		if err == nil {
			return int64(value)
		}
	}
	return fallback
}