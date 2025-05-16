package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // GANTI dengan env variable di production!

func GenerateToken(email string) (string, error) {
	// Buat claims (data yang akan disimpan dalam token)
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // token berlaku 24 jam
	}

	// Buat token dengan signing method dan claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tanda tangani token dan kembalikan
	return token.SignedString(jwtSecret)
}
