//go:build wireinject
// +build wireinject

package cms

import (
	"github.com/google/wire"

	"mocking_api/config/database"
)

func InitializeCms() routes {
	wire.Build(
		newRoutes,
		newController,
		newService,
		newRepository,
		database.InitSqlite,
	)
	return routes{}
}
