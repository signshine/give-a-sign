package service

import (
	"context"
	"errors"
	"time"

	"github.com/signshine/give-a-sign/api/pb"
	"github.com/signshine/give-a-sign/internal/user"
	"github.com/signshine/give-a-sign/internal/user/domain"
	"github.com/signshine/give-a-sign/internal/user/port"
	"github.com/signshine/give-a-sign/pkg/jwt"

	gojwt "github.com/golang-jwt/jwt/v5"
)

var (
	ErrUserOnCreate           = user.ErrUserOnCreate
	ErrUserCreationValidation = user.ErrUserCreationValidation
	ErrUserFilterValidation   = user.ErrUserFilterValidation
	ErrUserNotFound           = user.ErrUserNotFound

	ErrUserWrongPassword = errors.New("user wrong password")
)

type UserService struct {
	svc                   port.Service
	authSecret            string
	expMin, refreshExpMin uint
}

func NewUserService(svc port.Service, authSecret string, expMin, refreshExpMin uint) *UserService {
	return &UserService{
		svc:           svc,
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
	}
}

func (s *UserService) SignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	userID, err := s.svc.Create(ctx, &domain.User{
		Email:    domain.Email(req.Email),
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	access, refresh, err := s.CreateToken(userID)
	if err != nil {
		return nil, err
	}

	return &pb.UserSignUpResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *UserService) SignIn(ctx context.Context, req *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	user, err := s.svc.Get(ctx, &domain.UserFilter{
		Email: domain.Email(req.Email),
	})

	if err != nil {
		return nil, err
	}

	if !user.PasswordIsCorrect(req.Password) {
		return nil, ErrUserWrongPassword
	}

	access, refresh, err := s.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &pb.UserSignInResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, err
}

func (s *UserService) CreateToken(userID domain.UserID) (access, refresh string, err error) {
	var (
		accessExp  = time.Now().Add(time.Minute * time.Duration(s.expMin))
		refreshExp = time.Now().Add(time.Minute * time.Duration(s.expMin))
	)

	access, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(accessExp),
		},
		UserID: userID,
	})

	if err != nil {
		return
	}

	refresh, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(refreshExp),
		},
		UserID: userID,
	})

	if err != nil {
		return
	}

	return
}
