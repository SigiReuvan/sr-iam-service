package cache

import (
	"context"
	"errors"

	"github.com/SigiReuvan/iam-service/internal"
	"github.com/go-kit/log"
	"github.com/redis/go-redis/v9"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

type repository struct {
	cache  *redis.Client
	logger log.Logger
}

func New(cache *redis.Client, logger log.Logger) internal.Session {
	return &repository{
		cache:  cache,
		logger: logger,
	}
}

func (repo *repository) UserLogin(ctx context.Context, user internal.User) (string, error) {
	repo.createSession(ctx, user)
	err := repo.addAction(ctx, user, "LOGIN", "IAM")
	if err != nil {
		return "", err
	}
	return "", ErrNotImplemented
}

func (repo *repository) UserLogout(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (repo *repository) UserAuthenticate(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (repo *repository) UserAuthorize(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (repo *repository) RefreshToken(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}
