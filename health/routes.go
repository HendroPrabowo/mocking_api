package health

import (
	"github.com/go-chi/chi"
)

var APP_NAME = "mocking_api"

type routes struct {
	controller controller
}

func newRoutes(controller controller) routes {
	return routes{controller: controller}
}

func (routes routes) RegisterRoutes(r *chi.Mux) {
	r.Get("/status", routes.controller.HealthCheck())
}
