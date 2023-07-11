package transport

import (
	"context"

	"github.com/SigiReuvan/iam-service/internal"
	"github.com/SigiReuvan/iam-service/internal/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser       endpoint.Endpoint
	DeleteUser       endpoint.Endpoint
	UserLogin        endpoint.Endpoint
	UserLogout       endpoint.Endpoint
	UserAuthenticate endpoint.Endpoint
	UserAuthorize    endpoint.Endpoint
	UserProfile      endpoint.Endpoint
	RefreshToken     endpoint.Endpoint
	PasswordReset    endpoint.Endpoint
}

func makeCreateUserEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		user := internal.User{
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
		user := service.User(req.Username, req.Email, req.Password)
		token, err := svc.UserLogin(ctx, user)
		if err != nil {
			return userLoginResponse{"", err.Error()}, err
		}
		return userLoginResponse{Token: token, Err: ""}, err
	}
}

func makeUserLogoutEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userLogoutRequest)
		user := service.User(req.Username, req.Email, req.Password)
		if err := svc.UserLogout(ctx, user); err != nil {
			return userLogoutResponse{err.Error()}, err
		}
		return userLogoutResponse{Err: ""}, nil
	}
}

func makeRefreshTokenEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(refreshTokenRequest)
		user := service.User(req.Username, req.Email, req.Password)
		token, err := svc.RefreshToken(ctx, user)
		if err != nil {
			return refreshTokenResponse{"", err.Error()}, err
		}
		return refreshTokenResponse{Token: token, Err: ""}, err
	}
}

func makePasswordResetEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userLogoutRequest)
		user := service.User(req.Username, req.Email, req.Password)
		if err := svc.PasswordReset(ctx, user); err != nil {
			return passwordResetResponse{err.Error()}, err
		}
		return passwordResetResponse{Err: ""}, nil
	}
}

func makeUserAuthenticateEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(userAuthenticateRequest)
		user := service.User(req.Username, req.Email, req.Password)
		if err := svc.UserAuthenticate(ctx, user); err != nil {
			return userAuthenticateResponse{err.Error()}, err
		}
		return userAuthenticateResponse{Err: ""}, nil
	}
}

func makeUserAuthorizeEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(userAuthenticateRequest)
		user := service.User(req.Username, req.Email, req.Password)
		if err := svc.UserAuthorize(ctx, user); err != nil {
			return userAuthorizeResponse{err.Error()}, err
		}
		return userAuthorizeResponse{Err: ""}, nil
	}
}

func makeUserProfileEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(userProfileRequest)
		user := service.User(req.Username, req.Email, req.Password)
		result, err := svc.UserProfile(ctx, user)
		if err != nil {
			return userProfileResponse{Err: err.Error()}, err
		}
		return userProfileResponse{ID: result.ID, Username: result.Username, Email: result.Email, Err: ""}, nil
	}
}
