package health

import (
	"github.com/go-chi/chi"
	"github.com/hellofresh/health-go/v5"
)

var APP_NAME = "mocking_api"

func RegisterRoutes(r *chi.Mux) {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    APP_NAME,
		Version: "v1.0",
	}))

	r.Get("/status", h.HandlerFunc)
	r.Get("/", h.HandlerFunc)
}
