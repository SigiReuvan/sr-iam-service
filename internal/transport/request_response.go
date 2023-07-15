package transport

import "time"

type createUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}
type getUserRequest struct {
	ID string
}

type getUserResponse struct {
	ID        string    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Err       string    `json:"err,omitempty"`
}
type deleteUserRequest struct {
	ID string `json:"id"`
}

type deleteUserResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}

type userLoginRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type userLoginResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"`
}

type userLogoutResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}

type userAuthenticateResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}

type userAuthorizeResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}

type refreshTokenResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"`
}

type passwordResetRequest struct {
	Username         string `json:"username,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password"`
	NewPassword      string `json:"newPassword"`
	NewPasswordAgain string `json:"newPasswordAgain"`
}

type passwordResetResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}

type emptyRequest struct{}
