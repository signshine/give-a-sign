package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.New().String()
	return nil
}
