package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("iman")

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expired")
)

type TokenParams struct {
	Duration time.Duration
}

func CreateToken(tokenParams *TokenParams) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenParams.Duration).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func VerifyToken(token string) error {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return secretKey, nil
	}

	_, err := jwt.Parse(token, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return ErrExpiredToken
		}
		return ErrInvalidToken
	}

	return nil
}
