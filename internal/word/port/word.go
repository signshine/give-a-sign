package port

import (
	"context"

	"github.com/signshine/give-a-sign/internal/word/domain"
)

type Repo interface {
	CreateWord(ctx context.Context, word domain.Word) (domain.WordID, error)
	GetWord(ctx context.Context, filter domain.WordFilter) (*domain.Word, error)
	GetAllWords(ctx context.Context, page, pageSize int) ([]*domain.Word, error)
	DeleteWord(ctx context.Context, filter domain.WordFilter) error
}
