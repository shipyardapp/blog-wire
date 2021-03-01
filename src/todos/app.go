package todos

import (
	"context"
	"io"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/shipyardapp/blog-wire/src/infra/server"
)

type DB interface {
	io.Closer
}

type APM interface {
	io.Closer
}

type App struct {
	db  DB
	apm APM

	server *server.Server
}

func New(db DB, apm APM, server *server.Server) *App {
	return &App{
		db:     db,
		apm:    apm,
		server: server,
	}
}

func (a *App) Run(ctx context.Context) error {
	errChan := make(chan error)

	done := make(chan struct{})

	go func() {
		select {
		case <-ctx.Done():
			shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
			errChan <- a.server.Shutdown(shutdownCtx)
			cancel()

		case <-done:
		}

		close(errChan)
	}()

	go func() {
		errChan <- a.server.ListenAndServe()
		close(done)
	}()

	return <-errChan
}

func (a *App) Close() error {
	var result *multierror.Error

	for _, closer := range []io.Closer{a.db, a.apm} {
		if err := closer.Close(); err != nil {
			result = multierror.Append(result, err)
		}
	}

	return result.ErrorOrNil()
}
