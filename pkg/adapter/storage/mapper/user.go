package mapper

import (
	userDomain "github.com/signshine/give-a-sign/internal/user/domain"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/types"

	"gorm.io/gorm"
)

func UserDomain2Storage(user userDomain.User) *types.User {
	return &types.User{
		Model: gorm.Model{
			ID: uint(user.ID),
		},
		UUID:     user.UUID.String(),
		Email:    string(user.Email),
		Password: user.Password,
	}
}

func UserStorage2Domain(user types.User) *userDomain.User {
	return &userDomain.User{
		ID:       userDomain.UserID(user.ID),
		UUID:     userDomain.UserUUID([]byte(user.UUID)),
		Email:    userDomain.Email(user.Email),
		Password: user.Password,
	}
}
