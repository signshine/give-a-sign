package app

import (
	"context"

	"github.com/signshine/give-a-sign/config"
	"github.com/signshine/give-a-sign/internal/language"
	langPort "github.com/signshine/give-a-sign/internal/language/port"
	"github.com/signshine/give-a-sign/internal/user"
	userPort "github.com/signshine/give-a-sign/internal/user/port"
	"github.com/signshine/give-a-sign/internal/word"
	wordPort "github.com/signshine/give-a-sign/internal/word/port"
	"github.com/signshine/give-a-sign/pkg/adapter/storage"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/types"
	appCtx "github.com/signshine/give-a-sign/pkg/context"
	"github.com/signshine/give-a-sign/pkg/sqlite"

	"gorm.io/gorm"
)

type app struct {
	db          *gorm.DB
	cfg         config.Config
	wordService wordPort.Service
	userService userPort.Service
	langService langPort.Service
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}

func (a *app) setDB() error {
	db, err := sqlite.NewSQLiteGormConnection(sqlite.DBConnOptions{
		DBName: a.cfg.DB.Database,
	})

	if err != nil {
		return err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&types.User{},
		&types.Language{},
		&types.SignLanguage{},
		&types.Word{},
	)

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = userServiceWithDB(a.db)
		}
		return a.userService
	}

	return userServiceWithDB(db)
}

func userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db))
}

func (a *app) LanguageService(ctx context.Context) langPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.langService == nil {
			a.langService = languageServiceWithDB(a.db)
		}
		return a.langService
	}

	return languageServiceWithDB(db)
}

func languageServiceWithDB(db *gorm.DB) langPort.Service {
	return language.NewService(storage.NewLanguageRepo(db))
}

func (a *app) WordService(ctx context.Context) wordPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.wordService == nil {
			a.wordService = wordServiceWithDB(a.db)
		}
		return a.wordService
	}

	return wordServiceWithDB(db)
}

func wordServiceWithDB(db *gorm.DB) wordPort.Service {
	return word.NewService(storage.NewWordRepo(db))
}