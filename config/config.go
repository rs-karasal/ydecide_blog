package config

import "os"

var JwtSecretKey = []byte(getEnv("JWT_SECRET", "default-secret-key"))

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
