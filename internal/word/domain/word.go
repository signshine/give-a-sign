package domain

import (
	"github.com/google/uuid"
)

type (
	WordID   uint
	WordUUID = uuid.UUID
)

type Word struct {
	ID         WordID
	UUID       WordUUID
	Name       string
	LanguageID WordID
}

func (w *Word) IsValid() bool {
	return len(w.Name) > 0 && w.LanguageID > 0
}

type WordFilter struct {
	ID         WordID
	LanguageID WordID
	Name       string
}

func (f *WordFilter) IsValid() bool {
	return f.ID > 0 || f.LanguageID > 0 || len(f.Name) > 0
}
