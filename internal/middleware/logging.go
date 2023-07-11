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

func (mw logmw) CreateUser(ctx context.Context, user internal.User) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "createUser",
			"user", user.Username,
			"email", user.Email,
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.CreateUser(ctx, user)
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

func (mw logmw) UserLogin(ctx context.Context, user internal.User) (output string, err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "createUser",
			"request", requestString,
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.UserLogin(ctx, user)
	return output, err
}

func (mw logmw) UserLogout(ctx context.Context, user internal.User) (err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "createUser",
			"request", requestString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Service.UserLogout(ctx, user)
	return err
}

func (mw logmw) UserAuthenticate(ctx context.Context, user internal.User) (err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "createUser",
			"request", requestString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Service.UserAuthenticate(ctx, user)
	return err
}

func (mw logmw) UserAuthorize(ctx context.Context, user internal.User) (err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "createUser",
			"request", requestString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Service.UserAuthorize(ctx, user)
	return err
}

func (mw logmw) UserProfile(ctx context.Context, user internal.User) (output internal.User, err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "createUser",
			"request", requestString,
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Service.UserProfile(ctx, user)
	return output, err
}

func (mw logmw) RefreshToken(ctx context.Context, user internal.User) (output string, err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "createUser",
			"request", requestString,
			"ouput", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.RefreshToken(ctx, user)
	return output, err
}

func (mw logmw) PasswordReset(ctx context.Context, user internal.User) (err error) {
	defer func(begin time.Time) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		requestString := buf.String()
		_ = mw.logger.Log(
			"method", "createUser",
			"request", requestString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Service.PasswordReset(ctx, user)
	return err
}
