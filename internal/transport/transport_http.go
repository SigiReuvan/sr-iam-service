package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/SigiReuvan/iam-service/internal"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc internal.Service) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeErrorResponse),
		httptransport.ServerBefore(setCookies),
	}
	createUserHandler := httptransport.NewServer(
		makeCreateUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeResponse,
		options...,
	)
	deleteUserHandler := httptransport.NewServer(
		makeDeleteUserEndpoint(svc),
		decodeDeleteUserRequest,
		encodeResponse,
		options...,
	)
	userLoginHandler := httptransport.NewServer(
		makeUserLoginEndpoint(svc),
		decodeUserLoginRequest,
		encodeResponse,
		options...,
	)
	userLogoutHandler := httptransport.NewServer(
		makeUserLogoutEndpoint(svc),
		decodeUserLogoutRequest,
		encodeResponse,
		options...,
	)
	userAuthenticateHandler := httptransport.NewServer(
		makeUserAuthenticateEndpoint(svc),
		decodeUserAuthenticate,
		encodeResponse,
		options...,
	)
	userAuthorizeHandler := httptransport.NewServer(
		makeUserAuthorizeEndpoint(svc),
		decodeUserAuthorize,
		encodeResponse,
		options...,
	)

	userProfileHandler := httptransport.NewServer(
		makeUserProfileEndpoint(svc),
		decodeUserProfile,
		encodeResponse,
		options...,
	)
	refreshTokenHandler := httptransport.NewServer(
		makeRefreshTokenEndpoint(svc),
		decodeRefreshToken,
		encodeResponse,
		options...,
	)
	passwordResetHandler := httptransport.NewServer(
		makePasswordResetEndpoint(svc),
		decodePasswordReset,
		encodeResponse,
		options...,
	)
	r := mux.NewRouter()

	r.Methods("POST").Path("/v1/users").Handler(createUserHandler)
	r.Methods("DELETE").Path("/v1/users").Handler(deleteUserHandler)
	r.Methods("POST").Path("/v1/login").Handler(userLoginHandler)
	r.Methods("POST").Path("/v1/logout").Handler(userLogoutHandler)
	r.Methods("POST").Path("/v1/refresh-token").Handler(refreshTokenHandler)
	r.Methods("POST").Path("/v1/password-reset").Handler(passwordResetHandler)
	r.Methods("POST").Path("/v1/authenticate").Handler(userAuthenticateHandler)
	r.Methods("POST").Path("/v1/authorize").Handler(userAuthorizeHandler)
	r.Methods("GET").Path("/v1/profile").Handler(userProfileHandler)
	return r
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	errMessage := errFrom(err)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": errMessage,
	})
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request deleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUserLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request userLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUserLogoutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request userLogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUserAuthenticate(ctx context.Context, r *http.Request) (interface{}, error) {
	var request userAuthenticateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUserAuthorize(ctx context.Context, r *http.Request) (interface{}, error) {
	var request userAuthorizeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUserProfile(ctx context.Context, r *http.Request) (interface{}, error) {
	var request userProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeRefreshToken(ctx context.Context, r *http.Request) (interface{}, error) {
	var request refreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodePasswordReset(ctx context.Context, r *http.Request) (interface{}, error) {
	var request passwordResetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
