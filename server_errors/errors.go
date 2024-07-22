package server_errors

import (
	"encoding/json"
	"net/http"
)

type (
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
)

var (
	_ = NewError("Invalid http method", http.StatusNotFound)
	_ = NewError("Only GET allowed", http.StatusNotFound)
	_ = NewError("Invalid content type", http.StatusBadRequest)
)

func NewError(err string, code int) Error {
	e := Error{
		Status:  code,
		Message: err,
	}

	return e
}

func NewServerError(err string) Error {
	e := Error{
		Status:  http.StatusInternalServerError,
		Message: err,
	}
	return e
}

func NewBadRequest(err string) Error {
	e := Error{
		Status:  http.StatusBadRequest,
		Message: err,
	}
	return e
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) HttpCode() int {
	return e.Status
}

func (e *Error) JSON() []byte {
	b, _ := json.Marshal(e)
	return b
}

func ReplyWithError(w http.ResponseWriter, e Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HttpCode())
	_, _ = w.Write(e.JSON())
}

var _ = error(&Error{})
