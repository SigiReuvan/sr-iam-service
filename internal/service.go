package internal

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user User) (string, error)
	DeleteUser(ctx context.Context, user User) (string, error)
	UserLogin(ctx context.Context, user User) (string, error)
	UserLogout(ctx context.Context, user User) error
	UserAuthenticate(ctx context.Context, user User) error
	UserAuthorize(ctx context.Context, user User) error
	UserProfile(ctx context.Context, user User) (User, error)
	RefreshToken(ctx context.Context, user User) (string, error)
	PasswordReset(ctx context.Context, user User) error
}

type Service interface {
	CreateUser(ctx context.Context, user User) (string, error)
	DeleteUser(ctx context.Context, id string) (string, error)
	UserLogin(ctx context.Context, user User) (string, error)
	UserLogout(ctx context.Context, user User) error
	UserAuthenticate(ctx context.Context, user User) error
	UserAuthorize(ctx context.Context, user User) error
	UserProfile(ctx context.Context, user User) (User, error)
	RefreshToken(ctx context.Context, user User) (string, error)
	PasswordReset(ctx context.Context, user User) error
}
