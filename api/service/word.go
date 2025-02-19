package service

import (
	"context"

	"github.com/signshine/give-a-sign/internal/word/domain"
	"github.com/signshine/give-a-sign/internal/word/port"
)

type WordService struct {
	svc port.Service
}

func NewWordService(svc port.Service) *WordService {
	return &WordService{
		svc: svc,
	}
}

func (ws *WordService) CreateWord(ctx context.Context, word domain.Word) (domain.WordID, error) {
	id, err := ws.svc.CreateWord(ctx, word)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ws *WordService) GetWordById(ctx context.Context, id uint) (*domain.Word, error) {
	panic("unimplemented")
}

func (ws *WordService) GetVideoById(ctx context.Context, id uint) (*domain.Word, error) {
	panic("unimplemented")
}

func (ws *WordService) GetAllWords(ctx context.Context, page, pageSize int) ([]*domain.Word, error) {
	words, err := ws.svc.GetAllWords(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}
	return words, nil
}
