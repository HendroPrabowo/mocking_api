//+build wireinject

package dukcapil

import (
	"github.com/google/wire"
)

func InitializeDukcapil() (routes, error) {
	wire.Build(newRoutes, newController)
	return routes{}, nil
}
