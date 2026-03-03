package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret")

func GenerateToken(userID uint) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    fmt.Println("JWT_SECRET =", secret)

    jwtKey := []byte(secret)

    claims := jwt.MapClaims{
        "user_id": userID,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}
