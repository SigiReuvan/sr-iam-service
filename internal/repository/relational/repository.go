package relational

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

func (repo *repository) GetUser(ctx context.Context, user internal.User) (internal.User, error) {
	return internal.User{}, ErrNotImplemented
}

func (repo *repository) DeleteUser(ctx context.Context, user internal.User) (string, error) {
	result := repo.db.Delete(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return "success", nil
}

func (repo *repository) UserLogin(ctx context.Context, user internal.UserLoginForm) (internal.User, error) {
	if user.Username != "" {
		var u internal.User
		result := repo.db.Table("Users").Where("username = ?", user.Username).First(&u)
		if result.Error != nil {
			return internal.User{}, result.Error
		}
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
		if err != nil {
			return internal.User{}, ErrWrongPassword
		}
		return u, nil
	}

	if user.Email != "" {
		var u internal.User
		result := repo.db.Table("Users").Where("email = ?", user.Email).First(&u)
		if result.Error != nil {
			return internal.User{}, result.Error
		}
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
		if err != nil {
			return internal.User{}, ErrWrongPassword
		}
		return u, nil
	}
	return internal.User{}, ErrUsernameOrEmailMissing
}

func (repo *repository) PasswordReset(ctx context.Context, user internal.User) (string, error) {
	return "", ErrNotImplemented
}
