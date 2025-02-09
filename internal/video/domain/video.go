package domain

import (
	"errors"

	"github.com/google/uuid"
	langDomain "github.com/signshine/give-a-sign/internal/language/domain"
	wordDomain "github.com/signshine/give-a-sign/internal/word/domain"
)

var (
	ErrVideoInvalidUUID = errors.New("invalid video uuid")
	ErrVideoEmptyPath   = errors.New("video path is empty")
)

type (
	VideoID   uint
	VideoUUID = uuid.UUID
)

func ValidateVideoUUID(id VideoUUID) error {
	if err := uuid.Validate(id.String()); err != nil {
		return ErrVideoInvalidUUID
	}
	return nil
}

type Video struct {
	ID             VideoID
	UUID           VideoUUID
	Path           string
	WordID         wordDomain.WordID
	SignLanguageID langDomain.LanguageID
}

func (v *Video) Validate() error {
	if len(v.Path) == 0 {
		return ErrVideoEmptyPath
	}
	return nil
}

func (v *Video) IsValid() bool {
	return v.WordID > 0 && v.SignLanguageID > 0
}

type VideoFilter struct {
	ID             VideoID
	UUID           VideoUUID
	WordID         wordDomain.WordID
	SignLanguageID langDomain.SignLanguageID
}

func (f *VideoFilter) IsValid() bool {
	return f.ID > 0 ||
		f.WordID > 0 ||
		f.SignLanguageID > 0 ||
		f.UUID != VideoUUID{}
}
