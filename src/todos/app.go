package todos

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/shipyardapp/blog-wire/src/infra/server"
)

// DB is an interface that wraps the io.Closer interface.
//
// This interface is defined for a couple of reasons.
// 1. With Wire, we are not allowed to have multiple providers for a single type,
// and this type would conflict with APM.
// 2. This interface defines everything that this package needs of a database
// and nothing more.
// The application will close the connections to the database after shutdown,
// but this package does not need access to query the database, for example.
type DB interface {
	io.Closer
}

// APM is an interface the wraps the io.Closer interface.
// It exists for the same reasons as DB.
type APM interface {
	io.Closer
}

type App struct {
	db  DB
	apm APM

	server *server.Server
}

func newApp(db DB, apm APM, server *server.Server) *App {
	return &App{
		db:     db,
		apm:    apm,
		server: server,
	}
}

func (a *App) Run(ctx context.Context) error {
	errChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
		defer cancel()

		errChan <- a.server.Shutdown(shutdownCtx)
	}()

	go func() {
		defer wg.Done()

		errChan <- a.server.ListenAndServe()
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var result *multierror.Error
	for err := range errChan {
		if err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result.ErrorOrNil()

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
