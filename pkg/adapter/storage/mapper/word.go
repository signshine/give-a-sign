package mapper

import (
	"github.com/google/uuid"
	"github.com/signshine/give-a-sign/internal/word/domain"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/types"
	"gorm.io/gorm"
)

func WordDomain2storage(word domain.Word) *types.Word {
	return &types.Word{
		Model:       gorm.Model{
			ID: uint(word.ID),
		},
		UUID:        word.UUID.String(),
		Name:        word.Name,
		EnglishName: word.EnglishName,
		LanguageID:  uint(word.LanguageID),
	}
}

func WordStorage2Domain(word types.Word) *domain.Word {
	wordUUID, _ := uuid.Parse(word.UUID)
	return &domain.Word{
		ID:          domain.WordID(word.ID),
		UUID:        wordUUID,
		Name:        word.Name,
		EnglishName: word.EnglishName,
		LanguageID:  domain.WordID(word.LanguageID),
	}
}