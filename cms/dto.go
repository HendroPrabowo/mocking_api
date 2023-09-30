package cms

type MockDto struct {
	Id           int         `json:"id"`
	Name         string      `json:"name" validate:"required"`
	Method       string      `json:"method" validate:"required"`
	Path         string      `json:"path" validate:"required"`
	ResponseCode int         `json:"response_code" validate:"required"`
	Request      interface{} `json:"request"`
	Response     interface{} `json:"response" validate:"required"`
}

type MockQueryDto struct {
	Page         int
	Limit        int
	Name         string
	Method       string
	Path         string
	ResponseCode int
}

var mehtodAllowed = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"PATCH":  true,
	"DELETE": true,
}
