package mapper

import (
	"github.com/signshine/give-a-sign/internal/language/domain"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/types"
	"gorm.io/gorm"
)

func LanguageDomain2Storage(lang domain.Language) *types.Language {
	return &types.Language{
		Model: gorm.Model{ID: uint(lang.ID)},
		UUID:  lang.UUID.String(),
		Name:  lang.Name,
	}
}

func LanguageStorage2Domain(lang types.Language) *domain.Language {
	return &domain.Language{
		ID:   domain.LanguageID(lang.ID),
		UUID: domain.LanguageUUID([]byte(lang.UUID)),
		Name: lang.Name,
	}
}

func SignLanguageDomain2Storage(lang domain.SignLanguage) *types.SignLanguage {
	return &types.SignLanguage{
		Model: gorm.Model{ID: uint(lang.ID)},
		UUID:  lang.UUID.String(),
		Name:  lang.Name,
	}
}

func SignLanguageStorage2Domain(lang types.SignLanguage) *domain.SignLanguage {
	return &domain.SignLanguage{
		ID:   domain.SignLanguageID(lang.ID),
		UUID: domain.SignLanguageUUID([]byte(lang.UUID)),
		Name: lang.Name,
	}
}