package transport

import (
	"context"

	"github.com/SigiReuvan/iam-service/internal"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser       endpoint.Endpoint
	GetUser          endpoint.Endpoint
	DeleteUser       endpoint.Endpoint
	UserLogin        endpoint.Endpoint
	UserLogout       endpoint.Endpoint
	UserAuthenticate endpoint.Endpoint
	UserAuthorize    endpoint.Endpoint
	RefreshToken     endpoint.Endpoint
	PasswordReset    endpoint.Endpoint
}

func makeCreateUserEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		user := internal.UserCreateForm{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}
		result, err := svc.CreateUser(ctx, user)
		if err != nil {
			return createUserResponse{"", err.Error()}, err
		}
		return createUserResponse{Message: result, Err: ""}, err
	}
}

func makeGetUserEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getUserRequest)
		result, err := svc.GetUser(ctx, req.ID)
		if err != nil {
			return getUserResponse{Err: err.Error()}, err
		}
		return getUserResponse{
			ID:        result.ID,
			Username:  result.Username,
			Email:     result.Email,
			Role:      result.Role,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
			Err:       ""}, nil
	}
}

func makeDeleteUserEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		result, err := svc.DeleteUser(ctx, req.ID)
		if err != nil {
			return deleteUserResponse{"", err.Error()}, err
		}
		return deleteUserResponse{Message: result, Err: ""}, err
	}
}

func makeUserLoginEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userLoginRequest)
		u := internal.UserLoginForm{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}
		token, err := svc.UserLogin(ctx, u)
		if err != nil {
			return userLoginResponse{"", err.Error()}, err
		}
		return userLoginResponse{Token: token, Err: ""}, err
	}
}

func makeUserLogoutEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := svc.UserLogout(ctx)
		if err != nil {
			return userLogoutResponse{Message: result, Err: err.Error()}, err
		}
		return userLogoutResponse{Err: ""}, nil
	}
}

func makeRefreshTokenEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		token, err := svc.RefreshToken(ctx)
		if err != nil {
			return refreshTokenResponse{"", err.Error()}, err
		}
		return refreshTokenResponse{Token: token, Err: ""}, err
	}
}

func makePasswordResetEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(passwordResetRequest)
		u := internal.UserPasswordResetForm{
			Username:         req.Username,
			Email:            req.Email,
			OldPassword:      req.Password,
			NewPassword:      req.NewPassword,
			NewPasswordAgain: req.NewPasswordAgain,
		}
		result, err := svc.PasswordReset(ctx, u)
		if err != nil {
			return passwordResetResponse{Message: result, Err: err.Error()}, err
		}
		return passwordResetResponse{Message: result, Err: ""}, nil
	}
}

func makeUserAuthenticateEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		result, err := svc.UserAuthenticate(ctx)
		if err != nil {
			return userAuthenticateResponse{Message: result, Err: err.Error()}, err
		}
		return userAuthenticateResponse{Message: result, Err: ""}, nil
	}
}

func makeUserAuthorizeEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		result, err := svc.UserAuthorize(ctx)
		if err != nil {
			return userAuthorizeResponse{Message: result, Err: err.Error()}, err
		}
		return userAuthorizeResponse{Message: result, Err: ""}, nil
	}
}
