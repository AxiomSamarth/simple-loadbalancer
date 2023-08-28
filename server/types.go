package server

import (
	"net/http"
)

type Server interface {
	GetAddress() string
	IsAlive() bool
	Start() error
}

type SimpleServer struct {
	address string
	mux     *http.ServeMux
}
