package service

import (
	"context"

	"github.com/signshine/give-a-sign/api/pb"
	"github.com/signshine/give-a-sign/internal/word"
	"github.com/signshine/give-a-sign/internal/word/domain"
	"github.com/signshine/give-a-sign/internal/word/port"
	"github.com/signshine/give-a-sign/pkg/fp"
)

var (
	ErrWordOnCreate           = word.ErrWordOnCreate
	ErrWordCreationValidation = word.ErrWordCreationValidation
	ErrWordFilterValidation   = word.ErrWordFilterValidation
	ErrWordOnGet              = word.ErrWordOnGet
	ErrWordNotFound           = word.ErrWordNotFound
	ErrWordOnGetAll           = word.ErrWordOnGetAll
	ErrWordOnDelete           = word.ErrWordOnDelete
	ErrWordAlreadyExist       = word.ErrWordAlreadyExist
)

type WordService struct {
	svc port.Service
}

func NewWordService(svc port.Service) *WordService {
	return &WordService{
		svc: svc,
	}
}

func (ws *WordService) CreateWord(ctx context.Context, req *pb.CreateWordRequest) (*pb.CreateWordResponse, error) {
	wordId, err := ws.svc.CreateWord(ctx, domain.Word{
		Name:        req.Name,
		EnglishName: req.EnglishName,
		LanguageID:  domain.WordID(req.LanguageId),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateWordResponse{
		Word: &pb.Word{Id: uint64(wordId)},
	}, nil
}

func (ws *WordService) GetWord(ctx context.Context, req *pb.GetWordRequest) (*pb.GetWordResponse, error) {
	word, err := ws.svc.GetWord(ctx, domain.WordFilter{
		ID:         domain.WordID(req.Filter.Id),
		LanguageID: domain.WordID(req.Filter.LanguageId),
		Name:       req.Filter.Name,
	})

	if err != nil {
		return nil, err
	}

	return &pb.GetWordResponse{
		Word: WordDomain2PB(word),
	}, nil
}

func (ws *WordService) GetAllWords(ctx context.Context, req *pb.ListWordRequest) (*pb.ListWordResponse, error) {
	words, err := ws.svc.GetAllWords(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	return &pb.ListWordResponse{
		Words: fp.Map(words, func(w *domain.Word) *pb.Word {
			return WordDomain2PB(w)
		}),
		TotalCount: 0,
	}, nil
}

func (ws *WordService) DeleteWord(ctx context.Context, req *pb.DeleteWordRequest) (*pb.DeleteWordResponse, error) {
	err := ws.svc.DeleteWord(ctx, domain.WordFilter{
		ID:         domain.WordID(req.Filter.Id),
		LanguageID: domain.WordID(req.Filter.LanguageId),
		Name:       req.Filter.Name,
	})

	return &pb.DeleteWordResponse{
		Success: err == nil,
	}, err
}

func WordDomain2PB(word *domain.Word) *pb.Word {
	return &pb.Word{
		Id:          uint64(word.ID),
		Uuid:        word.UUID.String(),
		Name:        word.Name,
		EnglishName: word.EnglishName,
		LanguageId:  uint64(word.LanguageID),
	}
}
