package http

import (
	"context"

	"github.com/signshine/give-a-sign/api/service"
	"github.com/signshine/give-a-sign/app"
	"github.com/signshine/give-a-sign/config"
)

type serviceGetter[T any] func(context.Context) T

// user service transient instance handler
func UserServiceGetter(app app.App, cfg config.ServerConfig) serviceGetter[*service.UserService] {
	return func(ctx context.Context) *service.UserService {
		return service.NewUserService(app.UserService(ctx),
			cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)
	}
}

// word service transient instance handler
func WordServiceGetter(app app.App) serviceGetter[*service.WordService] {
	return func(ctx context.Context) *service.WordService {
		return service.NewWordService(app.WordService(ctx))
	}
}
