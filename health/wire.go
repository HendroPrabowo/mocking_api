//+build wireinject

package health

import "github.com/google/wire"

func InitializeHealthCheck() (routes, error) {
	wire.Build(newRoutes, newController)
	return routes{}, nil
}
