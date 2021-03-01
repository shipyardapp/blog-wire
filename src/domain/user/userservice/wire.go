//+build wireinject

package userservice

import "github.com/google/wire"

var Wired = wire.NewSet(
	NewConfig,
	New,
)
