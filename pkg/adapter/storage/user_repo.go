package storage

import (
	"context"
	"errors"

	userDomain "github.com/signshine/give-a-sign/internal/user/domain"
	userPort "github.com/signshine/give-a-sign/internal/user/port"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/mapper"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/types"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userPort.Repo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *userDomain.User) (userDomain.UserID, error) {
	u := mapper.UserDomain2Storage(*user)
	return userDomain.UserID(u.ID), r.db.Table("users").WithContext(ctx).Create(u).Error
}
func (r *userRepo) Get(ctx context.Context, filter *userDomain.UserFilter) (*userDomain.User, error) {
	var user types.User

	q := r.db.Table("users").Debug().WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if filter.UUID != userDomain.NilUUID {
		q = q.Where("uuid = ?", filter.UUID.String())
	}

	if filter.Email.IsValid() {
		q = q.Where("email = ?", filter.Email.String())
	}

	err := q.First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}
