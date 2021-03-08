package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/go-multierror"
	"github.com/shipyardapp/blog-wire/src/todos"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app, err := todos.New(os.Getenv, log.Default())
	if err != nil {
		exit(err, 1)
	}

	exitCode := 0

	err = app.Run(ctx)
	if err != nil {
		exitCode = 2
	}

	if errClose := app.Close(); errClose != nil {
		err = multierror.Append(err, errClose)
		exitCode = 3
	}

	if err != nil {
		exit(err, exitCode)
	}
}

func exit(err error, exitCode int) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(exitCode)
}
