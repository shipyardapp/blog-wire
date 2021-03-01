//+build wireinject

package todorepo

import "github.com/google/wire"

var Wired = wire.NewSet(
	New,
)
