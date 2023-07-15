package transport

import (
	"net/http"

	"github.com/SigiReuvan/iam-service/internal/repository/relational"
)

// // TODO: Should be cookies even used?
// type contextKey string

// const (
// 	cookieKey contextKey = "iam-cookie"
// )

func codeFrom(err error) int {
	switch err {
	case relational.ErrNotImplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusInternalServerError
	}

}

func errFrom(err error) string {
	switch err.Error() {
	case relational.ErrNotUniqueEmail.Error():
		return relational.ErrNotUniqueEmail.Error()
	case relational.ErrNotUniqueUsername.Error():
		return relational.ErrNotUniqueUsername.Error()
	case relational.ErrWrongPassword.Error():
		return relational.ErrWrongPassword.Error()
	default:
		return err.Error()
	}
}

// func setCookies(ctx context.Context, r *http.Request) context.Context {
// 	s := securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
// 	value := "some-value"

// 	if encoded, err := s.Encode("iam-cookie", value); err == nil {
// 		cookie := &http.Cookie{
// 			Name:     "iam-cookie",
// 			Value:    encoded,
// 			Path:     "/v1/login",
// 			Expires:  time.Now().Add(24 * time.Hour),
// 			HttpOnly: true,
// 			SameSite: http.SameSiteStrictMode,
// 		}
// 		r.AddCookie(cookie)
// 		ctx = context.WithValue(ctx, cookieKey, cookie)

// 		return ctx
// 	}
// 	return nil

// }
