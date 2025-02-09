package word

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/signshine/give-a-sign/internal/word/domain"
	"github.com/signshine/give-a-sign/internal/word/port"
)

var (
	ErrWordOnCreate           = errors.New("error on creating new Word")
	ErrWordCreationValidation = errors.New("word validation failed")
	ErrWordFilterValidation   = errors.New("word filter validation failed")
	ErrWordOnGet              = errors.New("error on getting word")
	ErrWordNotFound           = errors.New("word not found")
	ErrWordOnGetAll           = errors.New("error on getting all words")
	ErrWordOnDelete           = errors.New("error on deleting word")

	ErrPaginationNegativePage     = errors.New("pagination error: page cannot be negative")
	ErrPaginationNegativePageSize = errors.New("pagination error: page size cannot be negative")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo: repo}
}

func (s *service) CreateWord(ctx context.Context, word domain.Word) (domain.WordID, error) {
	if !word.IsValid() {
		return 0, ErrWordCreationValidation
	}

	id, err := s.repo.CreateWord(ctx, word)
	if err != nil {
		log.Printf("")
		return 0, fmt.Errorf("%w, %w", ErrWordOnCreate, err)
	}

	return id, nil
}

func (s *service) GetWord(ctx context.Context, filter domain.WordFilter) (*domain.Word, error) {
	if !filter.IsValid() {
		return nil, ErrWordFilterValidation
	}

	word, err := s.repo.GetWord(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrWordOnGet, err)
	}

	if word == nil {
		return nil, ErrWordNotFound
	}

	return word, nil
}

func (s *service) GetAllWords(ctx context.Context, page, pageSize int) ([]*domain.Word, error) {
	if page < 0 {
		return nil, ErrPaginationNegativePage
	}
	if pageSize < 0 {
		return nil, ErrPaginationNegativePageSize
	}

	words, err := s.repo.GetAllWords(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrWordOnGetAll, err)
	}

	if words == nil {
		words = []*domain.Word{}
	}

	return words, nil
}

func (s *service) DeleteWord(ctx context.Context, filter domain.WordFilter) error {
	if !filter.IsValid() {
		return ErrWordFilterValidation
	}

	err := s.repo.DeleteWord(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w, %w", ErrWordOnDelete, err)
	}

	return nil
}
