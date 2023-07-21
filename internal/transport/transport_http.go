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
	}
	createUserHandler := httptransport.NewServer(
		makeCreateUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeResponse,
		options...,
	)
	getUserHandler := httptransport.NewServer(
		makeGetUserEndpoint(svc),
		decodeGetUserRequest,
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
		decodeEmtpyRequest,
		encodeResponse,
		options...,
	)
	userAuthenticateHandler := httptransport.NewServer(
		makeUserAuthenticateEndpoint(svc),
		decodeEmtpyRequest,
		encodeResponse,
		options...,
	)
	userAuthorizeHandler := httptransport.NewServer(
		makeUserAuthorizeEndpoint(svc),
		decodeEmtpyRequest,
		encodeResponse,
		options...,
	)
	refreshTokenHandler := httptransport.NewServer(
		makeRefreshTokenEndpoint(svc),
		decodeEmtpyRequest,
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
	r.Methods("GET").Path("/v1/users/{id}").Handler(getUserHandler)
	r.Methods("DELETE").Path("/v1/users").Handler(deleteUserHandler)
	r.Methods("POST").Path("/v1/login").Handler(userLoginHandler)
	r.Methods("POST").Path("/v1/logout").Handler(userLogoutHandler)
	r.Methods("POST").Path("/v1/refresh-token").Handler(refreshTokenHandler)
	r.Methods("POST").Path("/v1/password-reset").Handler(passwordResetHandler)
	r.Methods("POST").Path("/v1/authenticate").Handler(userAuthenticateHandler)
	r.Methods("POST").Path("/v1/authorize").Handler(userAuthorizeHandler)
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

func decodeGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	request := getUserRequest{
		ID: id,
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

func decodeEmtpyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request emptyRequest
	return request, nil
}
