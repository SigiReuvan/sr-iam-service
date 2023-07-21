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

type ActiveSession struct {
	UID            string
	Role           string
	Creation       time.Time
	Expiration     time.Time
	SessionActions []SessionAction
}

type SessionAction struct {
	Action       string
	Resource     string
	LastActivity time.Time
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
