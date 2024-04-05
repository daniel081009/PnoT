package auth

import (
	db_auth "PnoT/db/auth"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func ValidateJWT(tokenString string) (string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("Invalid token")
	}

	username := claims["username"].(string)
	return username, nil
}

func CreateUser(username string, password string) error {
	if db_auth.ExistsUser(username) {
		return errors.New("user already exists")
	}
	e := db_auth.CreateUser(username, password)
	if e != nil {
		return e
	}
	return nil
}

func LoginUser(username string, password string) (string, error) {
	err := db_auth.LoginUser(username, password)
	if err != nil {
		return "", err
	}
	token, err := GenerateJWT(username)
	if err != nil {
		return "", err
	}
	return token, nil
}
