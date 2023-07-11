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
)

type service struct {
	repository internal.Repository
	logger     log.Logger
}

func NewService(rep internal.Repository, logger log.Logger) internal.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func User(username string, email string, password string) internal.User {
	return internal.User{
		Username: username,
		Email:    email,
		Password: password,
	}
}

func (s *service) CreateUser(ctx context.Context, user internal.User) (string, error) {
	user, err := defaultUser(user)
	if err != nil {
		return "", err
	}
	result, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) (string, error) {
	user := internal.User{
		ID: id,
	}
	result, err := s.repository.DeleteUser(ctx, user)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *service) UserLogin(ctx context.Context, user internal.User) (string, error) {
	token, err := s.repository.UserLogin(ctx, user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) UserLogout(ctx context.Context, user internal.User) error {
	if err := s.repository.UserLogout(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *service) UserAuthenticate(ctx context.Context, user internal.User) error {
	if err := s.repository.UserAuthenticate(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *service) UserAuthorize(ctx context.Context, user internal.User) error {
	if err := s.repository.UserAuthorize(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *service) UserProfile(ctx context.Context, user internal.User) (internal.User, error) {
	result, err := s.repository.UserProfile(ctx, user)
	if err != nil {
		return internal.User{}, err
	}
	return result, nil
}

func (s *service) RefreshToken(ctx context.Context, user internal.User) (string, error) {
	result, err := s.repository.RefreshToken(ctx, user)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *service) PasswordReset(ctx context.Context, user internal.User) error {
	if err := s.repository.PasswordReset(ctx, user); err != nil {
		return err
	}
	return nil
}
