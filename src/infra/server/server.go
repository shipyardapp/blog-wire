package server

import (
	"context"
	"net/http"

	"github.com/shipyardapp/blog-wire/src/config"
)

type Config struct {
	ListenAddress string
}

func NewConfig(enver config.Enver) (Config, error) {
	// TODO

	return Config{
		ListenAddress: ":0",
	}, nil
}

type Server struct {
	server *http.Server
}

func New(config Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    config.ListenAddress,
			Handler: handler,
		},
	}
}

func (s *Server) ListenAndServe() error {
	err := s.server.ListenAndServe()
	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
