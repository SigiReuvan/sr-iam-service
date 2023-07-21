package service

import (
	"context"
	"errors"

	"github.com/SigiReuvan/iam-service/internal"
	"github.com/go-kit/log"
)

var (
	ErrBadPassword    = errors.New("password is longer than 36 characters")
	ErrService        = errors.New("unable to handle request")
	ErrNotImplemented = errors.New("not implemented")
	ErrUnableToLogin  = errors.New("unable to login")
)

type service struct {
	repository internal.Repository
	session    internal.Session
	logger     log.Logger
}

func NewService(rep internal.Repository, session internal.Session, logger log.Logger) internal.Service {
	return &service{
		repository: rep,
		session:    session,
		logger:     logger,
	}
}

func (svc *service) CreateUser(ctx context.Context, user internal.UserCreateForm) (string, error) {
	u, err := defaultUser(user)
	if err != nil {
		return "", err
	}
	result, err := svc.repository.CreateUser(ctx, u)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (svc *service) GetUser(ctx context.Context, id string) (internal.User, error) {
	u := internal.User{
		ID: id,
	}
	result, err := svc.repository.GetUser(ctx, u)
	if err != nil {
		return internal.User{}, err
	}
	return result, nil
}

func (svc *service) DeleteUser(ctx context.Context, id string) (string, error) {
	u := internal.User{
		ID: id,
	}
	result, err := svc.repository.DeleteUser(ctx, u)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (svc *service) UserLogin(ctx context.Context, user internal.UserLoginForm) (string, error) {
	u, err := svc.repository.UserLogin(ctx, user)
	if err != nil {
		return "", ErrUnableToLogin
	}

	token, err := svc.session.UserLogin(ctx, u)
	if err != nil {
		return "", ErrUnableToLogin
	}
	return token, nil

}

func (svc *service) UserLogout(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (svc *service) UserAuthenticate(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (svc *service) UserAuthorize(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (svc *service) RefreshToken(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (svc *service) PasswordReset(ctx context.Context, user internal.UserPasswordResetForm) (string, error) {
	return "", ErrNotImplemented
}
