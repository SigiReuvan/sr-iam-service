package service

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/SigiReuvan/iam-service/internal"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func defaultUser(form internal.UserCreateForm) (user internal.User, err error) {
	uuid, _ := uuid.NewV4()
	id := uuid.String()

	verificationCode, err := generateVerificationCode()
	if err != nil {
		return internal.User{}, err
	}

	if len([]byte(user.Password)) >= 36 {
		return internal.User{}, ErrBadPassword
	}

	hashedPassword, err := hashPassword(form.Password)
	if err != nil {
		return internal.User{}, err
	}

	return internal.User{
		ID:               id,
		Username:         form.Username,
		Email:            form.Email,
		Password:         hashedPassword,
		Role:             "user",
		VerificationCode: verificationCode,
		Verified:         false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, err
}

func generateVerificationCode() (int, error) {
	code, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return 0, err
	}
	code.Add(code, big.NewInt(100000))
	return int(code.Int64()), nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
