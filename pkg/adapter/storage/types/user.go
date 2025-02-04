package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string 
	Email    string `gorm:"uniqueIndex"`
	Password string
}
