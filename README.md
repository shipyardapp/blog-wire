# blog-wire

This is the repository for the code talked about in our XXXX blog post.

## Layout

The layout of the application is as follows:
* `src/config` - Package for configuration types and helpers.
* `src/domain` - Houses all of our business logic types and functionality.
* `src/domain/apm` - Provides an interface for an Application Peformance Monitoring server. We don’t actually set one up in the application, but it makes for a more interesting Wire setup.
* `src/domain/todo` - Provides the Todo type and interfaces for interacting with them.
* `src/domain/user` - Provides the User type and interfaces for interacting with them.
* `src/infra` - Provides real world implementations of some domain interfaces. Infra packages are the actual infrastructure connections in the application. Subpackages provide an http server, a sql database connection along with repo implementations, and a fake apm implementation.
* `src/todos` - The top level application package.
* `main.go` - The main package that we compile.

### Wire

The top level wire.Build call in in `src/todos/wire.go`.
