package cms

import "github.com/go-chi/chi"

type routes struct {
	controller controller
}

func newRoutes(controller controller) routes {
	return routes{
		controller: controller,
	}
}

func (routes routes) RegisterRoutes(r *chi.Mux) {
	r.Route("/api/v1/mock", func(r chi.Router) {
		r.Get("/", routes.controller.GetMock)
		r.Post("/", routes.controller.AddMock)
		r.Put("/", routes.controller.UpdateMock)
		r.Delete("/", routes.controller.DeleteMock)
	})
}

func (routes routes) RegisterServiceRoutes(r *chi.Mux) {
	r.Get("/*", routes.controller.HandleMock)
	r.Post("/*", routes.controller.HandleMock)
	r.Put("/*", routes.controller.HandleMock)
	r.Patch("/*", routes.controller.HandleMock)
	r.Delete("/*", routes.controller.HandleMock)
}
