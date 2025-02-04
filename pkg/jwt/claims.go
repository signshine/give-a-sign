package jwt

import (
	"github.com/signshine/give-a-sign/internal/user/domain"

	gojwt "github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	gojwt.RegisteredClaims
	UserID domain.UserID
}
