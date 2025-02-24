package storage

import (
	"context"
	"errors"

	"github.com/signshine/give-a-sign/internal/language"
	"github.com/signshine/give-a-sign/internal/language/domain"
	"github.com/signshine/give-a-sign/internal/language/port"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/mapper"
	"github.com/signshine/give-a-sign/pkg/adapter/storage/types"
	"github.com/signshine/give-a-sign/pkg/fp"
	"gorm.io/gorm"
)

const (
	TableLanguage     = "languages"
	TableSignLanguage = "sign_languages"
)

type languageRepo struct {
	db *gorm.DB
}

func NewLanguageRepo(db *gorm.DB) port.Repo {
	return &languageRepo{
		db: db,
	}
}

func (r *languageRepo) CreateLanguage(ctx context.Context, domainLang domain.Language) (domain.LanguageID, error) {
	lang := mapper.LanguageDomain2Storage(domainLang)
	q := r.db.Table(TableLanguage).WithContext(ctx)
	var err error

	// Restore the soft-deleted user
	var deletedLang types.Language
	err = q.Unscoped().Where("name = ?", lang.Name).First(&deletedLang).Error // Find the soft-deleted record
	if err == nil { 
		err = q.Unscoped().Model(&deletedLang).Update("DeletedAt", nil).Error // Restore by setting DeletedAt to nil
		return domain.LanguageID(deletedLang.ID), err
	}

	err = q.Where("name = ?", lang.Name).First(&types.Language{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = q.Create(lang).Error
		return domain.LanguageID(lang.ID), err
	}

	return 0, language.ErrLanguageAlreadyExist
}

func (r *languageRepo) CreateSignLanguage(ctx context.Context, domainLang domain.SignLanguage) (domain.SignLanguageID, error) {
	lang := mapper.SignLanguageDomain2Storage(domainLang)
	q := r.db.Table(TableSignLanguage).WithContext(ctx)
	var err error

	// Restore the soft-deleted user
	var deletedLang types.SignLanguage
	err = q.Unscoped().Where("name = ?", lang.Name).First(&deletedLang).Error // Find the soft-deleted record
	if err == nil { 
		err = q.Unscoped().Model(&deletedLang).Update("DeletedAt", nil).Error // Restore by setting DeletedAt to nil
		return domain.SignLanguageID(deletedLang.ID), err
	}

	err = q.Where("name = ?", lang.Name).First(&types.SignLanguage{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = q.Create(lang).Error
		return domain.SignLanguageID(lang.ID), err
	}

	return 0, language.ErrSignLanguageAlreadyExist
}

func (r *languageRepo) GetLanguage(ctx context.Context, filter domain.LanguageFilter) (*domain.Language, error) {
	var lang types.Language

	q := r.db.Table(TableLanguage).WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if filter.UUID != domain.NilUUID {
		q = q.Where("uuid = ?", filter.UUID)
	}

	if len(filter.Name) > 0 {
		q = q.Where("name = ?", filter.Name)
	}

	err := q.First(&lang).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.LanguageStorage2Domain(lang), nil
}

func (r *languageRepo) GetSignLanguage(ctx context.Context, filter domain.SignLanguageFilter) (*domain.SignLanguage, error) {
	var lang types.SignLanguage

	q := r.db.Table(TableSignLanguage).WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if filter.UUID != domain.NilUUID {
		q = q.Where("uuid = ?", filter.UUID)
	}

	if len(filter.Name) > 0 {
		q = q.Where("name = ?", filter.Name)
	}

	err := q.First(&lang).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mapper.SignLanguageStorage2Domain(lang), nil
}

func (r *languageRepo) GetAllLanguage(ctx context.Context, page, pageSize int) ([]*domain.Language, error) {
	var langs []types.Language

	offset := (page - 1) * pageSize

	err := r.db.WithContext(ctx).Table(TableLanguage).Offset(offset).Limit(pageSize).Find(&langs).Error
	if err != nil {
		return nil, err
	}

	return fp.Map(langs, func(l types.Language) *domain.Language {
		return mapper.LanguageStorage2Domain(l)
	}), nil
}

func (r *languageRepo) GetAllSignLanguage(ctx context.Context, page, pageSize int) ([]*domain.SignLanguage, error) {
	var langs []types.SignLanguage

	offset := (page - 1) * pageSize

	q := r.db.WithContext(ctx).Table(TableSignLanguage).Offset(offset).Limit(pageSize)

	err := q.Find(langs).Error
	if err != nil {
		return nil, err
	}

	return fp.Map(langs, func(l types.SignLanguage) *domain.SignLanguage {
		return mapper.SignLanguageStorage2Domain(l)
	}), nil
}

func (r *languageRepo) DeleteLanguage(ctx context.Context, filter domain.LanguageFilter) error {
	q := r.db.Table(TableLanguage).WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if filter.UUID != domain.NilUUID {
		q = q.Where("uuid = ?", filter.UUID)
	}

	if len(filter.Name) > 0 {
		q = q.Where("name = ?", filter.Name)
	}

	return q.Delete(&types.Language{}).Error
}

func (r *languageRepo) DeleteSignLanguage(ctx context.Context, filter domain.SignLanguageFilter) error {
	q := r.db.Table(TableSignLanguage).WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if filter.UUID != domain.NilUUID {
		q = q.Where("uuid = ?", filter.UUID)
	}

	if len(filter.Name) > 0 {
		q = q.Where("name = ?", filter.Name)
	}

	err := q.Delete(&types.SignLanguage{}).Error
	if err != nil {
		return err
	}
	return nil
}
