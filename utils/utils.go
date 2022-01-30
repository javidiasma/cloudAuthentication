package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var jwtKey = "4$m@J@v!d!"

func CreateTokens(username, password string) (string, error) {
	var err error
	expirationTime := time.Now().Add(72 * time.Hour)

	claims := &Claims{
		Password: password,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := myToken.SignedString([]byte(jwtKey))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string) (bool, *Claims) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, &Claims{}
		}
		return false, &Claims{}
	}
	if !tkn.Valid {
		return false, &Claims{}
	}
	return true, claims
}
func GetToken(ctx *gin.Context) (*Claims, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 {
		return &Claims{}, errors.New("unauthorized1")
	}
	validated, userClaim := ValidateToken(tokenParts[1])
	if !validated {
		return &Claims{}, errors.New("unauthorized2")
	}

	return userClaim, nil
}
