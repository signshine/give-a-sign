package storage

import (
	"context"
	"errors"

	"github.com/signshine/give-a-sign/internal/word"
	"github.com/signshine/give-a-sign/internal/word/domain"
	"github.com/signshine/give-a-sign/internal/word/port"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/mapper"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/types"
	"github.com/signshine/give-a-sign/pkg/fp"
	"gorm.io/gorm"
)

const (
	TableWord = "words"
)

type wordRepo struct {
	db *gorm.DB
}

func NewWordRepo(db *gorm.DB) port.Repo {
	return &wordRepo{db: db}
}

func (r *wordRepo) CreateWord(ctx context.Context, domainWord domain.Word) (domain.WordID, error) {
	storageWord := mapper.WordDomain2storage(domainWord)
	
	wordID, err := r.restoreWord(ctx, storageWord)
	if err == nil {
		return domain.WordID(wordID), nil
	}

	err = r.db.WithContext(ctx).Table(TableWord).Create(storageWord).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return 0, word.ErrWordAlreadyExist
	}

	return domain.WordID(storageWord.ID), err
}

// restoreWord Restore the soft-deleted word
func (r *wordRepo) restoreWord(ctx context.Context, w *types.Word) (uint, error) {
	db := r.db.WithContext(ctx)
	var deletedWord types.Word
	err := db.Table(TableWord).Unscoped().Where("name = ?", w.Name).First(&deletedWord).Error // Find the soft-deleted record
	if err != nil {
		return 0, err
	}
	err = db.Table(TableWord).Unscoped().Model(&deletedWord).Update("DeletedAt", nil).Error // Restore by setting DeletedAt to nil
	return deletedWord.ID, err
}

func (r *wordRepo) GetWord(ctx context.Context, filter domain.WordFilter) (*domain.Word, error) {
	var word types.Word

	q := r.db.WithContext(ctx).Table(TableWord)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if len(filter.Name) > 0 {
		q = q.Where("name = ?", filter.Name)
	}

	if filter.LanguageID > 0 {
		q = q.Where("language_id = ?", filter.LanguageID)
	}

	err := q.First(&word).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.WordStorage2Domain(word), nil
}

func (r *wordRepo) GetAllWords(ctx context.Context, page, pageSize int) ([]*domain.Word, error) {
	var words []types.Word
	q := r.db.WithContext(ctx).Table(TableWord)

	offset := (page - 1) * pageSize

	err := q.Offset(offset).Limit(pageSize).Find(&words).Error
	if err != nil {
		return nil, err
	}

	return fp.Map(words, func(w types.Word) *domain.Word {
		return mapper.WordStorage2Domain(w)
	}), nil
}

func (r *wordRepo) DeleteWord(ctx context.Context, filter domain.WordFilter) error {
	q := r.db.WithContext(ctx).Table(TableWord)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if len(filter.Name) > 0 {
		q = q.Where("name = ?", filter.Name)
	}

	return q.Delete(&types.Word{}).Error
}
