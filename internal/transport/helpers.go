package transport

import (
	"context"
	"net/http"
	"time"

	"github.com/SigiReuvan/iam-service/internal/repository"
	"github.com/SigiReuvan/iam-service/internal/service"
	"github.com/gorilla/securecookie"
)

// TODO: Should be cookies even used?
type contextKey string

const (
	cookieKey contextKey = "iam-cookie"
)

func codeFrom(err error) int {
	switch err {
	case repository.ErrNotImplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusInternalServerError
	}

}

func errFrom(err error) string {
	switch err.Error() {
	case repository.ErrNotUniqueEmail.Error():
		return repository.ErrNotUniqueEmail.Error()
	case repository.ErrNotUniqueUsername.Error():
		return repository.ErrNotUniqueUsername.Error()
	case repository.ErrNotUniqueUsername.Error():
		return service.ErrBadPassword.Error()
	default:
		return err.Error()
	}
}

func setCookies(ctx context.Context, r *http.Request) context.Context {
	s := securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
	value := "some-value"

	if encoded, err := s.Encode("iam-cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:     "iam-cookie",
			Value:    encoded,
			Path:     "/v1/login",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		r.AddCookie(cookie)
		ctx = context.WithValue(ctx, cookieKey, cookie)

		return ctx
	}
	return nil

}
