package client_server

import (
	"net/http"
	c "pafaul/overcomplicated-image-compressor/controllers"
	m "pafaul/overcomplicated-image-compressor/middleware"
)

func NewServer(addr string) *http.Server {
	s := &http.Server{
		Addr:    addr,
		Handler: newServeMux(),
	}

	return s
}

func newServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", m.HttpMethod(c.ServeFile("./public/file-upload.html"), http.MethodGet))
	mux.Handle("/request/", m.HttpMethod(m.MultipartForm(http.HandlerFunc(requestHandler)), http.MethodPost))
	mux.Handle("/health", m.HttpMethod(http.HandlerFunc(c.HealthStatus), http.MethodGet))
	return mux
}
