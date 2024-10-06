package config

import "os"

// TODO: add .env with real secret key
var JwtSecretKey = []byte(getEnv("JWT_SECRET", "default-secret-key"))

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
