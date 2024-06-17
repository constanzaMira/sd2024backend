package service

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"time"
)

type JWTService interface {
	GenerateToken(email string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	claims := &jwtClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "washup",
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Expira en 72 horas
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("washup"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, nil
		}

		return []byte("washup"), nil
	})
}

func hashPassword(plainPassword string) string {
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 12)
	if err != nil {
		return "error"
	}
	return string(bcryptHash)
}

func ValidatePassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}