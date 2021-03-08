//+build wireinject

package todos

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/wire"
	"github.com/shipyardapp/blog-wire/src/config"
	"github.com/shipyardapp/blog-wire/src/domain/apm"
	"github.com/shipyardapp/blog-wire/src/domain/todo"
	"github.com/shipyardapp/blog-wire/src/domain/todo/todoservice"
	"github.com/shipyardapp/blog-wire/src/domain/user"
	"github.com/shipyardapp/blog-wire/src/domain/user/userservice"
	"github.com/shipyardapp/blog-wire/src/infra/server"
	"github.com/shipyardapp/blog-wire/src/infra/somerealapm"
	"github.com/shipyardapp/blog-wire/src/infra/sqldb"
	"github.com/shipyardapp/blog-wire/src/infra/todorepo"
	"github.com/shipyardapp/blog-wire/src/infra/userrepo"
	"github.com/shipyardapp/blog-wire/src/todos/handler/api"
)

// New is the App injector.
// Note that we still have the newApp function in this package to provide an App.
//
// Earlier interations of this project used New for the App provider and
// InitializeApp as the injector, but we've updated the names to fit with our
// New constructor convention.
func New(enver config.Enver, logger *log.Logger) (*App, error) {
	wire.Build(

		// sqldb to DB.
		sqldb.NewConfig,
		sqldb.New,
		wire.Bind(new(DB), new(*sql.DB)),

		// Our imposter APM.
		somerealapm.NewConfig,
		somerealapm.New,
		wire.Bind(new(apm.APM), new(*somerealapm.APM)),
		wire.Bind(new(APM), new(*somerealapm.APM)),

		// User Repo.
		userrepo.Wired,
		wire.Bind(new(user.Repo), new(*userrepo.Repo)),

		// User Service.
		userservice.Wired,
		wire.Bind(new(user.Service), new(*userservice.Service)),

		// Todo Repo.
		todorepo.Wired,
		wire.Bind(new(todo.Repo), new(*todorepo.Repo)),

		// Todo Service.
		todoservice.Wired,
		wire.Bind(new(todo.Service), new(*todoservice.Service)),

		// API as our http.Handler.
		api.Wired,
		wire.Bind(new(http.Handler), new(*api.API)),

		// Http server.
		server.Wired,

		// This package - the application.
		newApp,
	)

	// Requirement for compilation.
	return nil, nil
}
