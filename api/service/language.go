package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/signshine/give-a-sign/api/pb"
	"github.com/signshine/give-a-sign/internal/language"
	"github.com/signshine/give-a-sign/internal/language/domain"
	"github.com/signshine/give-a-sign/internal/language/port"
	"github.com/signshine/give-a-sign/pkg/fp"
)

var (
	ErrLanguageOnCreate         = language.ErrLanguageOnCreate
	ErrLanguageAlreadyExist     = language.ErrLanguageAlreadyExist
	ErrLanguageOnGet            = language.ErrLanguageOnGet
	ErrLanguageNotFound         = language.ErrLanguageNotFound
	ErrLanguageFilterValidation = language.ErrLanguageFilterValidation

	ErrSignLanguageOnCreate         = language.ErrSignLanguageOnCreate
	ErrSignLanguageAlreadyExist     = language.ErrSignLanguageAlreadyExist
	ErrSignLanguageOnGet            = language.ErrSignLanguageOnGet
	ErrSignLanguageNotFound         = language.ErrSignLanguageNotFound
	ErrSignLanguageFilterValidation = language.ErrSignLanguageFilterValidation

	ErrPaginationNegativePage     = language.ErrPaginationNegativePage
	ErrPaginationNegativePagesize = language.ErrPaginationNegativePageSize
)

type LanguageService struct {
	svc port.Service
}

func NewLanguageService(svc port.Service) *LanguageService {
	return &LanguageService{svc: svc}
}

func (s *LanguageService) CreateLanguage(ctx context.Context, req *pb.CreateLanguageRequest) (*pb.CreateLanguageResponse, error) {
	langId, err := s.svc.CreateLanguage(ctx, domain.Language{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateLanguageResponse{
		Language: &pb.Language{Id: uint64(langId)},
		Error:    &pb.Error{},
	}, nil
}

func (s *LanguageService) CreateSignLanguage(ctx context.Context, req *pb.CreateSignLanguageRequest) (*pb.CreateSignLanguageResponse, error) {
	langId, err := s.svc.CreateSignLanguage(ctx, domain.SignLanguage{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateSignLanguageResponse{
		Language: &pb.SignLanguage{Id: uint64(langId)},
		Error:    &pb.Error{},
	}, nil
}

func (s *LanguageService) GetLanguage(ctx context.Context, req *pb.GetLanguageRequest) (*pb.GetLanguageResponse, error) {
	lang, err := s.svc.GetLanguage(ctx, *LanguageFilterPB2Domain(req.Filter))

	if err != nil {
		return nil, err
	}

	return &pb.GetLanguageResponse{
		Language: LanguageDomain2PB(lang),
		Error:    &pb.Error{},
	}, nil
}

func (s *LanguageService) GetSignLanguage(ctx context.Context, req *pb.GetSignLanguageRequest) (*pb.GetSignLanguageResponse, error) {
	lang, err := s.svc.GetSignLanguage(ctx, *SignLanguageFilterPB2Domain(req.Filter))

	if err != nil {
		return nil, err
	}

	return &pb.GetSignLanguageResponse{
		Language: SignLanguageDomain2PB(lang),
		Error:    &pb.Error{},
	}, nil
}

func (s *LanguageService) GetAllLanguage(ctx context.Context, req *pb.ListLanguagesRequest) (*pb.ListLanguagesResponse, error) {
	langs, err := s.svc.GetAllLanguage(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	return &pb.ListLanguagesResponse{
		Languages: fp.Map(langs, func(l *domain.Language) *pb.Language {
			return LanguageDomain2PB(l)
		}),
		TotalCount: 0,
		Error:      &pb.Error{},
	}, nil
}

func (s *LanguageService) GetAllSignLanguage(ctx context.Context, req *pb.ListSignLanguagesRequest) (*pb.ListSignLanguagesResponse, error) {
	langs, err := s.svc.GetAllSignLanguage(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	return &pb.ListSignLanguagesResponse{
		Languages: fp.Map(langs, func(l *domain.SignLanguage) *pb.SignLanguage {
			return SignLanguageDomain2PB(l)
		}),
		TotalCount: 0,
		Error:      &pb.Error{},
	}, nil
}

func (s *LanguageService) DeleteLanguage(ctx context.Context, req *pb.DeleteLanguageRequest) (*pb.DeleteLanguageResponse, error) {
	err := s.svc.DeleteLanguage(ctx, *LanguageFilterPB2Domain(req.Filter))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteLanguageResponse{
		Success: true,
		Error:   &pb.Error{},
	}, nil
}

func (s *LanguageService) DeleteSignLanguage(ctx context.Context, req *pb.DeleteSignLanguageRequest) (*pb.DeleteSignLanguageResponse, error) {
	err := s.svc.DeleteSignLanguage(ctx, *SignLanguageFilterPB2Domain(req.Filter))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteSignLanguageResponse{
		Success: true,
		Error:   &pb.Error{},
	}, nil
}

func LanguageDomain2PB(lang *domain.Language) *pb.Language {
	return &pb.Language{
		Id:   uint64(lang.ID),
		Uuid: lang.UUID.String(),
		Name: lang.Name,
	}
}

func LanguageFilterPB2Domain(filter *pb.LanguageFilter) *domain.LanguageFilter {
	filterUUID, _ := uuid.Parse(filter.Uuid)
	return &domain.LanguageFilter{
		ID:   domain.LanguageID(filter.Id),
		UUID: filterUUID,
		Name: filter.Name,
	}
}

func SignLanguageFilterPB2Domain(filter *pb.SignLanguageFilter) *domain.SignLanguageFilter {
	filterUUID, _ := uuid.Parse(filter.Uuid)
	return &domain.SignLanguageFilter{
		ID:   domain.SignLanguageID(filter.Id),
		UUID: filterUUID,
		Name: filter.Name,
	}
}

func SignLanguageDomain2PB(lang *domain.SignLanguage) *pb.SignLanguage {
	return &pb.SignLanguage{
		Id:   uint64(lang.ID),
		Uuid: lang.UUID.String(),
		Name: lang.Name,
	}
}
