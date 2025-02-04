package app

import (
	"context"

	"github.com/signshine/give-a-sign/config"
	userPort "github.com/signshine/give-a-sign/internal/user/port"
	wordPort "github.com/signshine/give-a-sign/internal/word/port"

	"gorm.io/gorm"
)

type App interface {
	UserService(ctx context.Context) userPort.Service
	WordService(ctx context.Context) wordPort.Service
	Config() config.Config
	DB() *gorm.DB
}
