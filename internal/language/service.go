package language

import (
	"context"
	"errors"
	"fmt"

	"github.com/signshine/give-a-sign/internal/language/domain"
	"github.com/signshine/give-a-sign/internal/language/port"
	// appCtx "github.com/signshine/give-a-sign/pkg/context"
)

var (
	ErrLanguageOnCreate           = errors.New("error on creating new language")
	ErrLanguageCreationValidation = errors.New("language validation failed")
	ErrLanguageFilterValidation   = errors.New("language filter validation failed")
	ErrLanguageOnGet              = errors.New("error on getting language")
	ErrLanguageNotFound           = errors.New("language not found")
	ErrLanguageOnGetAll           = errors.New("error on getting all language")
	ErrLanguageOnDelete           = errors.New("error on deleting language")
	ErrLanguageAlreadyExist       = errors.New("language already exists")

	ErrSignLanguageOnCreate           = errors.New("error on creating new sign language")
	ErrSignLanguageCreationValidation = errors.New("sign language validation failed")
	ErrSignLanguageFilterValidation   = errors.New("sign language filter validation failed")
	ErrSignLanguageOnGet              = errors.New("error on getting sign language")
	ErrSignLanguageNotFound           = errors.New("sign language not found")
	ErrSignLanguageOnGetAll           = errors.New("error on getting all sign language")
	ErrSignLanguageOnDelete           = errors.New("error on deleting sign language")
	ErrSignLanguageAlreadyExist       = errors.New("sign language already exists")

	ErrPaginationNegativePage     = errors.New("pagination error: page cannot be negative")
	ErrPaginationNegativePageSize = errors.New("pagination error: page size cannot be negative")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo: repo}
}

func (s *service) CreateLanguage(ctx context.Context, lang domain.Language) (domain.LanguageID, error) {
	// logger := appCtx.GetLogger(ctx)

	if err := lang.Validate(); err != nil {
		return 0, fmt.Errorf("%w: %w", ErrLanguageCreationValidation, err)
	}

	id, err := s.repo.CreateLanguage(ctx, lang)
	if err != nil {
		// logger.Debug(err.Error())
		if errors.Is(err, ErrLanguageAlreadyExist) {
			return 0, ErrLanguageAlreadyExist
		}
		return 0, ErrLanguageOnCreate
	}

	return id, nil
}

func (s *service) CreateSignLanguage(ctx context.Context, lang domain.SignLanguage) (domain.SignLanguageID, error) {
	if err := lang.Validate(); err != nil {
		return 0, fmt.Errorf("%w: %w", ErrSignLanguageCreationValidation, err)
	}

	id, err := s.repo.CreateSignLanguage(ctx, lang)
	if err != nil {
		// log
		if errors.Is(err, ErrSignLanguageAlreadyExist) {
			return 0, ErrSignLanguageAlreadyExist
		}
		return 0, ErrSignLanguageOnCreate
	}

	return id, nil
}

func (s *service) GetLanguage(ctx context.Context, filter domain.LanguageFilter) (*domain.Language, error) {
	if !filter.IsValid() {
		return nil, ErrLanguageFilterValidation
	}

	language, err := s.repo.GetLanguage(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrLanguageOnGet, err)
	}

	if language == nil {
		return nil, ErrLanguageNotFound
	}

	return language, nil
}

func (s *service) GetSignLanguage(ctx context.Context, filter domain.SignLanguageFilter) (*domain.SignLanguage, error) {
	if !filter.IsValid() {
		return nil, ErrSignLanguageFilterValidation
	}

	signLanguage, err := s.repo.GetSignLanguage(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrSignLanguageOnGet, err)
	}

	if signLanguage == nil {
		return nil, ErrSignLanguageNotFound
	}

	return signLanguage, nil
}

func (s *service) GetAllLanguage(ctx context.Context, page, pageSize int) ([]*domain.Language, error) {
	if page < 0 {
		return nil, ErrPaginationNegativePage
	}
	if pageSize < 0 {
		return nil, ErrPaginationNegativePageSize
	}

	languages, err := s.repo.GetAllLanguage(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrLanguageOnGetAll, err)
	}

	if languages == nil {
		languages = []*domain.Language{}
	}

	return languages, nil
}

func (s *service) GetAllSignLanguage(ctx context.Context, page, pageSize int) ([]*domain.SignLanguage, error) {
	if page < 0 {
		return nil, ErrPaginationNegativePage
	}
	if pageSize < 0 {
		return nil, ErrPaginationNegativePageSize
	}

	signLanguages, err := s.repo.GetAllSignLanguage(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrSignLanguageOnGetAll, err)
	}

	if signLanguages == nil {
		signLanguages = []*domain.SignLanguage{}
	}

	return signLanguages, nil
}

func (s *service) DeleteLanguage(ctx context.Context, filter domain.LanguageFilter) error {
	if !filter.IsValid() {
		return ErrLanguageFilterValidation
	}

	err := s.repo.DeleteLanguage(ctx, filter)
	if err != nil {
		return ErrLanguageOnDelete
	}

	return nil
}

func (s *service) DeleteSignLanguage(ctx context.Context, filter domain.SignLanguageFilter) error {
	if !filter.IsValid() {
		return ErrSignLanguageFilterValidation
	}

	err := s.repo.DeleteSignLanguage(ctx, filter)
	if err != nil {
		return ErrSignLanguageOnDelete
	}

	return nil
}
