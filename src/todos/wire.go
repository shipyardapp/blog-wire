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

func InitializeApp(enver config.Enver, logger *log.Logger) (*App, error) {
	wire.Build(

		// sqldb to DB
		sqldb.NewConfig,
		sqldb.New,
		wire.Bind(new(DB), new(*sql.DB)),

		somerealapm.NewConfig,
		somerealapm.New,
		wire.Bind(new(apm.APM), new(*somerealapm.APM)),
		wire.Bind(new(APM), new(*somerealapm.APM)),

		userrepo.Wired,
		wire.Bind(new(user.Repo), new(*userrepo.Repo)),

		userservice.Wired,
		wire.Bind(new(user.Service), new(*userservice.Service)),

		todorepo.Wired,
		wire.Bind(new(todo.Repo), new(*todorepo.Repo)),

		todoservice.Wired,
		wire.Bind(new(todo.Service), new(*todoservice.Service)),

		api.Wired,
		wire.Bind(new(http.Handler), new(*api.API)),

		server.Wired,

		// This package.
		New,
	)

	return nil, nil
}
