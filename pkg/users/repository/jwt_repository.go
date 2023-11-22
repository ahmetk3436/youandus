package repository

import (
	"errors"
	"strconv"
	"time"
	"youandus/pkg/users/middleware"

	"github.com/golang-jwt/jwt"
)

func CreateToken(exp time.Duration, username string, userId uint) (*string, error) {

	privateKey := []byte("youanduseventplannerv1")
	// Create a token with the private key
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(exp).Unix(),
	}
	// convert uint to string
	userIdStr := strconv.Itoa(int(userId))

	// create JWT with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JwtCustom{
		UserID:         userIdStr,
		UserName:       username,
		StandardClaims: claims,
	})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return nil, errors.New(err.Error() + " jwt imzalanırken problem oluştu !")
	}
	return &tokenString, nil
}
