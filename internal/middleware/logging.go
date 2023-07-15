package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/SigiReuvan/iam-service/internal"
	"github.com/go-kit/log"
)

type logmw struct {
	logger log.Logger
	internal.Service
}

// TODO: Find a way to properly propogate errors
func NewLoggingMiddleware(logger log.Logger, next internal.Service) logmw {
	return logmw{logger, next}
}

func (mw logmw) CreateUser(ctx context.Context, form internal.UserCreateForm) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "createUser",
			"user", form.Username,
			"email", form.Email,
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.CreateUser(ctx, form)
	return output, err
}

func (mw logmw) GetUser(ctx context.Context, id string) (output internal.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "getUser",
			"id", id,
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Service.GetUser(ctx, id)
	return output, err
}

func (mw logmw) DeleteUser(ctx context.Context, id string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "deleteUser",
			"id", id,
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.DeleteUser(ctx, id)
	return output, err
}

func (mw logmw) UserLogin(ctx context.Context, form internal.UserLoginForm) (output string, err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(form)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "UserLogin",
			"request", requestString,
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.UserLogin(ctx, form)
	return output, err
}

func (mw logmw) UserLogout(ctx context.Context) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "UserLogout",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.UserLogout(ctx)
	return output, err
}

func (mw logmw) UserAuthenticate(ctx context.Context) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "UserAuthenticate",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.UserAuthenticate(ctx)
	return output, err
}

func (mw logmw) UserAuthorize(ctx context.Context) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "UserAuthorize",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.UserAuthorize(ctx)
	return output, err
}

func (mw logmw) RefreshToken(ctx context.Context) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "refreshToken",
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.RefreshToken(ctx)
	return output, err
}

func (mw logmw) PasswordReset(ctx context.Context, form internal.UserPasswordResetForm) (output string, err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(form)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "passwordReset",
			"request", requestString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.PasswordReset(ctx, form)
	return output, err
}
