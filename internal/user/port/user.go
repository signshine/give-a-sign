package port

import (
	"context"

	"github.com/signshine/give-a-sign/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user *domain.User) (domain.UserID, error)
	Get(ctx context.Context, filter *domain.UserFilter) (*domain.User, error)
}
