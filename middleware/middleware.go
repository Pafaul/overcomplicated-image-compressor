package middleware

import (
	"fmt"
	"net/http"
	"pafaul/overcomplicated-image-compressor/server_errors"
	"strings"
)

var multipartFormData = "multipart/form-data"

type Handler func(w http.ResponseWriter, r *http.Request)

func HttpMethod(next http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			msg := fmt.Sprintf("Method used: %s, expected: %s", r.Method, method)
			server_errors.ReplyWithError(w, server_errors.NewServerError(msg))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MultipartForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ct := r.Header.Get("Content-Type"); !strings.Contains(ct, multipartFormData) {
			msg := fmt.Sprintf("received content-type: %s, expected: %s", ct, multipartFormData)
			server_errors.ReplyWithError(w, server_errors.NewBadRequest(msg))
			return
		}

		next.ServeHTTP(w, r)
	})
}
