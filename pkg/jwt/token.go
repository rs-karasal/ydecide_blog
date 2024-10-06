package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs-karasal/ydecide_blog/config"
)

func GenerateToken(username string, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(duration).Unix()

	return token.SignedString(config.JwtSecretKey)
}
