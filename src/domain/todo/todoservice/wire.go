//+build wireinject

package todoservice

import "github.com/google/wire"

var Wired = wire.NewSet(
	NewConfig,
	New,
)
