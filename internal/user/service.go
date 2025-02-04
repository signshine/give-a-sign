package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/signshine/give-a-sign/internal/user/domain"
	"github.com/signshine/give-a-sign/internal/user/port"
)

var (
	ErrUserOnCreate           = errors.New("error on creating new user")
	ErrUserCreationValidation = errors.New("user validation failed")
	ErrUserFilterValidation   = errors.New("user filter validation failed")
	ErrUserNotFound           = errors.New("user not found")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, user *domain.User) (domain.UserID, error) {
	if err := user.Validate(); err != nil {
		return 0, fmt.Errorf("%w, %w", ErrUserCreationValidation, err)
	}

	user.Password = domain.NewPassword(user.Password)

	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		log.Printf("error on creation new user: %v", err.Error())
		return 0, ErrUserOnCreate
	}

	return userID, nil
}

func (s *service) Get(ctx context.Context, filter *domain.UserFilter) (*domain.User, error) {
	if !filter.IsValid() {
		return nil, ErrUserFilterValidation
	}

	user, err := s.repo.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
