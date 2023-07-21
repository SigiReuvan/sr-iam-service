package transport

import (
	"net/http"

	"github.com/SigiReuvan/iam-service/internal/repository/relational"
)

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
