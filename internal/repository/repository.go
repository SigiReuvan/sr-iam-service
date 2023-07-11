package repository

import (
	"context"
	"errors"

	"github.com/SigiReuvan/iam-service/internal"
	"github.com/go-kit/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrRepository             = errors.New("unable to handle request")
	ErrNotImplemented         = errors.New("not implemented")
	ErrNotUniqueEmail         = errors.New("mail already exists")
	ErrNotUniqueUsername      = errors.New("user already exists")
	ErrUsernameOrEmailMissing = errors.New("username or email missing")
	ErrWrongPassword          = errors.New("wrong password")
)

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func New(db *gorm.DB, logger log.Logger) internal.Repository {
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "postgres"),
	}
}

func (repo *repository) CreateUser(ctx context.Context, user internal.User) (string, error) {
	unique, err := repo.checkUniqueness("username", user.Username)
	if unique == "failed" && err != nil {
		return "", err
	}
	if unique == "failed" && err == nil {
		return "", ErrNotUniqueUsername
	}

	unique, err = repo.checkUniqueness("email", user.Email)
	if unique == "failed" && err != nil {
		return "", err
	}
	if unique == "failed" && err == nil {
		return "", ErrNotUniqueEmail
	}

	result := repo.db.Create(&user)
	if result.Error != nil {
		return "", err
	}
	return "success", nil
}

func (repo *repository) DeleteUser(ctx context.Context, user internal.User) (string, error) {
	result := repo.db.Delete(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return "success", nil
}

func (repo *repository) UserLogin(ctx context.Context, user internal.User) (string, error) {
	if user.Username != "" {
		var search internal.User
		result := repo.db.Table("Users").Where("username = ?", user.Username).First(&search)
		if result.Error != nil {
			return "failed", result.Error
		}
		err := bcrypt.CompareHashAndPassword([]byte(search.Password), []byte(user.Password))
		if err != nil {
			return "failed", ErrWrongPassword
		}
		return "success", nil
	}

	if user.Email != "" {
		var search internal.User
		result := repo.db.Table("Users").Where("email = ?", user.Email).First(&search)
		if result.Error != nil {
			return "failed", result.Error
		}
		err := bcrypt.CompareHashAndPassword([]byte(search.Password), []byte(user.Password))
		if err != nil {
			return "failed", ErrWrongPassword
		}
		return "success", nil
	}
	return "failed", ErrUsernameOrEmailMissing
}

func (repo *repository) UserLogout(ctx context.Context, user internal.User) error {
	return ErrNotImplemented
}

func (repo *repository) UserAuthenticate(ctx context.Context, user internal.User) error {
	return ErrNotImplemented
}

func (repo *repository) UserAuthorize(ctx context.Context, user internal.User) error {
	return ErrNotImplemented
}

func (repo *repository) UserProfile(ctx context.Context, user internal.User) (internal.User, error) {
	return internal.User{}, ErrNotImplemented
}

func (repo *repository) RefreshToken(ctx context.Context, user internal.User) (string, error) {
	return "", ErrNotImplemented
}

func (repo *repository) PasswordReset(ctx context.Context, user internal.User) error {
	return ErrNotImplemented
}
