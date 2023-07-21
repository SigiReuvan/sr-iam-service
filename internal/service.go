package internal

import (
	"context"
)

type Session interface {
	UserLogin(ctx context.Context, user User) (string, error)
	UserLogout(ctx context.Context) (string, error)
	UserAuthenticate(ctx context.Context) (string, error)
	UserAuthorize(ctx context.Context) (string, error)
	RefreshToken(ctx context.Context) (string, error)
}

type Repository interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetUser(ctx context.Context, user User) (User, error)
	DeleteUser(ctx context.Context, user User) (string, error)
	UserLogin(ctx context.Context, user UserLoginForm) (User, error)
	PasswordReset(ctx context.Context, user User) (string, error)
}

type Service interface {
	CreateUser(ctx context.Context, form UserCreateForm) (string, error)
	GetUser(ctx context.Context, username string) (User, error)
	DeleteUser(ctx context.Context, id string) (string, error)
	UserLogin(ctx context.Context, user UserLoginForm) (string, error)
	UserLogout(ctx context.Context) (string, error)
	UserAuthenticate(ctx context.Context) (string, error)
	UserAuthorize(ctx context.Context) (string, error)
	RefreshToken(ctx context.Context) (string, error)
	PasswordReset(ctx context.Context, form UserPasswordResetForm) (string, error)
}
