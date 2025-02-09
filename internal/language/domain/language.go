package domain

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

var (
	ErrInvalidLanguageUUID     = errors.New("invalid language uuid")
	ErrInvalidSignLanguageUUID = errors.New("invalid sign language uuid")
	ErrLanguageWithoutName     = errors.New("error language without name")
	ErrSignLanguageWithoutName = errors.New("error sign language without name")
	ErrInvalidLanguageName     = errors.New("invalid language name")
	ErrInvalidSignLanguageName = errors.New("invalid sign language name")
)

var languageRegex = regexp.MustCompile(`^[\p{L}]+$`)

type (
	LanguageID   uint
	LanguageUUID = uuid.UUID
)

func ValidateLanguageUUID(id LanguageUUID) error {
	if err := uuid.Validate(id.String()); err != nil {
		return ErrInvalidLanguageUUID
	}
	return nil
}

type Language struct {
	ID   LanguageID
	UUID LanguageUUID
	Name string
}

func (l *Language) Validate() error {
	if len(l.Name) == 0 {
		return ErrLanguageWithoutName
	}
	if !languageRegex.MatchString(l.Name) {
		return ErrInvalidLanguageName
	}
	return nil
}

type LanguageFilter struct {
	ID   LanguageID
	UUID LanguageUUID
	Name string
}

func (f *LanguageFilter) IsValid() bool {
	return f.ID > 0 || f.UUID != LanguageUUID{} || languageRegex.MatchString(f.Name)
}

type (
	SignLanguageID   uint
	SignLanguageUUID = uuid.UUID
)

func ValidateSignLanguageUUID(id LanguageUUID) error {
	if err := uuid.Validate(id.String()); err != nil {
		return ErrInvalidSignLanguageUUID
	}
	return nil
}

type SignLanguage struct {
	ID   SignLanguageID
	UUID SignLanguageUUID
	Name string
}

func (l *SignLanguage) Validate() error {
	if len(l.Name) == 0 {
		return ErrSignLanguageWithoutName
	}
	if !languageRegex.MatchString(l.Name) {
		return ErrInvalidSignLanguageName
	}
	return nil
}

type SignLanguageFilter struct {
	ID   SignLanguageID
	UUID SignLanguageUUID
	Name string
}

func (f *SignLanguageFilter) IsValid() bool {
	return f.ID > 0 || f.UUID != LanguageUUID{} || languageRegex.MatchString(f.Name)
}
