package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	UUID        string
	Name        string `gorm:"uniqueIndex"`
	EnglishName string 
	LanguageID  uint
}

func (w *Word) BeforeCreate(tx *gorm.DB) error {
	w.UUID = uuid.New().String()
	return nil
}
