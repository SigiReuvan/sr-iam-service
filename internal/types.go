package internal

import "time"

type User struct {
	ID               string
	Username         string
	Email            string
	Password         string
	Role             string
	VerificationCode int
	Verified         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserCreateForm struct {
	Username string
	Email    string
	Password string
}

type UserLoginForm struct {
	Username string
	Email    string
	Password string
}

type UserRefreshTokenForm struct {
	id           string
	RefreshToken string
}

type UserPasswordResetForm struct {
	Username         string
	Email            string
	OldPassword      string
	NewPassword      string
	NewPasswordAgain string
}
