package dukcapil

import "github.com/go-chi/chi"

type routes struct {
	controller controller
}

func newRoutes(controller controller) routes {
	return routes{controller: controller}
}

func (routes routes) RegisterRoutes(r *chi.Mux) {
	r.Post("/dukcapil/get_json/{id}/CALL_VERIFY_BY_ELEMEN", routes.controller.DukcapilIdentityVerify)
}


