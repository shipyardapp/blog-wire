//+build wireinject

package userrepo

import "github.com/google/wire"

var Wired = wire.NewSet(
	New,
)
