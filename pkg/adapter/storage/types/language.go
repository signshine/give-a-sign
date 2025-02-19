package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	UUID string
	Name string `gorm:"uniqueIndex"`
}

func (l *Language) BeforeCreate(tx *gorm.DB) error {
	l.UUID = uuid.New().String()
	return nil
}

type SignLanguage struct {
	gorm.Model
	UUID string
	Name string `gorm:"uniqueIndex"`
}

func (l *SignLanguage) BeforeCreate(tx *gorm.DB) error {
	l.UUID = uuid.New().String()
	return nil
}
