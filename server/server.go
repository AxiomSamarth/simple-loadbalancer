package server

import (
	"fmt"
	"net/http"
)

func (s *SimpleServer) GetAddress() string {
	return s.address
}

func (s *SimpleServer) IsAlive() bool {
	// Simple check for liveness, just return true
	return true
}

func NewSimpleServer(address string) Server {
	server := &SimpleServer{
		address: address,
		mux:     http.NewServeMux(),
	}
	server.registerHandlers()
	return server
}

func (s *SimpleServer) registerHandlers() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
}

func (s *SimpleServer) Start() error {
	fmt.Printf("\nServer started on %s", s.GetAddress())
	return http.ListenAndServe(s.GetAddress(), s.mux)
}
