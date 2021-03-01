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
	var err error
	c := Config{}

	if c.ListenAddress, err = config.RequiredString(enver, "BLOGWIRE_SERVER_LISTEN_ADDRESS"); err != nil {
		return Config{}, err
	}

	return c, nil
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
