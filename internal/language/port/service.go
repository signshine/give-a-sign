package port

import (
	"context"

	"github.com/signshine/give-a-sign/internal/language/domain"
)

type Service interface {
	CreateLanguage(ctx context.Context, lang domain.Language) (domain.LanguageID, error)
	CreateSignLanguage(ctx context.Context, lang domain.SignLanguage) (domain.SignLanguageID, error)

	GetLanguage(ctx context.Context, filter domain.LanguageFilter) (*domain.Language, error)
	GetSignLanguage(ctx context.Context, filter domain.SignLanguageFilter) (*domain.SignLanguage, error)

	GetAllLanguage(ctx context.Context, page, pageSize int) ([]*domain.Language, error)
	GetAllSignLanguage(ctx context.Context, page, pageSize int) ([]*domain.SignLanguage, error)

	DeleteLanguage(ctx context.Context, filter domain.LanguageFilter) error
	DeleteSignLanguage(ctx context.Context, filter domain.SignLanguageFilter) error
}
