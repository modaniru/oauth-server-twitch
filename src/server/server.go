package server

import (
	"net/http"
	"time"
)

type server struct {
	server *http.Server
}

func NewServer(handler http.Handler, port string) *server {
	return &server{
		server: &http.Server{
			Handler:        handler,
			Addr:           ":" + port,
			MaxHeaderBytes: 1 << 20, //1 MG
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *server) Run() {
	s.server.ListenAndServe()
}
