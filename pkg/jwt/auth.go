package jwt

import (
	"errors"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

var (
	ErrNilToken     = errors.New("invalid token (nil)")
	ErrInvalidToken = errors.New("token is not valid")
)

func CreateToken(secret []byte, claims *UserClaims) (string, error) {
	return jwt2.NewWithClaims(jwt2.SigningMethodHS512, claims).SignedString(secret)
}

func ParseToken(tokenString string, secret []byte) (*UserClaims, error) {
	token, err := jwt2.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt2.Token) (interface{}, error) {
		return secret, nil
	})

	if token == nil {
		return nil, ErrNilToken
	}

	var claim *UserClaims
	if token.Claims != nil {
		cc, ok := token.Claims.(*UserClaims)
		if ok {
			claim = cc
		}
	}

	if err != nil {
		return claim, err
	}

	if !token.Valid {
		return claim, ErrInvalidToken
	}

	return claim, nil
}
