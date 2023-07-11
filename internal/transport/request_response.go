package transport

type createUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
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

type userLogoutRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type userLogoutResponse struct {
	Err string `json:"err,omitempty"`
}

type userAuthenticateRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type userAuthenticateResponse struct {
	Err string `json:"err,omitempty"`
}

type userAuthorizeRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type userAuthorizeResponse struct {
	Err string `json:"err,omitempty"`
}

type userProfileRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type userProfileResponse struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Err      string `json:"err,omitempty"`
}

type refreshTokenRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type refreshTokenResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"`
}

type passwordResetRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type passwordResetResponse struct {
	Err string `json:"err,omitempty"`
}
