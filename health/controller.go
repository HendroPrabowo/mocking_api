package health

import (
	"net/http"

	"github.com/hellofresh/health-go/v5"
)

type controller struct {
}

func newController() controller {
	return controller{}
}

func (c controller) HealthCheck() http.HandlerFunc {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    APP_NAME,
		Version: "v1.0",
	}))

	return h.HandlerFunc
}
