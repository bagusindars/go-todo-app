package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`

	jwt.RegisteredClaims
}

func GetJwtSecret() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		return "", errors.New("jwt secret has not been set")
	}

	return jwtSecret, nil
}

func GenerateToken(userId uint, email string, name string) (string, error) {
	tokenExpired := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserId: userId,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpired),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// generate token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret, err := GetJwtSecret()

	if err != nil {
		return "", err
	}

	signedAccessToken, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New("Error signed token : " + err.Error())
	}

	return signedAccessToken, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	jwtSecret, err := GetJwtSecret()

	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid and return the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
