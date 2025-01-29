package cms

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type MockResponseDto struct {
	Mock         []MockDto `json:"mock"`
	Page         int       `json:"page"`
	TotalElement int       `json:"total_element"`
}

type MockDto struct {
	Id           uint                   `json:"id"`
	Name         string                 `json:"name" validate:"required"`
	Method       string                 `json:"method" validate:"required"`
	Path         string                 `json:"path" validate:"required"`
	ResponseCode int                    `json:"response_code" validate:"required"`
	Request      map[string]interface{} `json:"request"`
	Response     map[string]interface{} `json:"response" validate:"required"`
}

func (dto *MockDto) fromEntity(mock Mock) {
	dto.Id = mock.ID
	dto.Name = mock.Name
	dto.Method = mock.Method
	dto.Path = mock.Path
	dto.ResponseCode = mock.ResponseCode
	if mock.Request != "" {
		if err := json.Unmarshal([]byte(mock.Request), &dto.Request); err != nil {
			log.Warning("failed to unmarshal mock request")
		}
	}

	if mock.Response != "" {
		if err := json.Unmarshal([]byte(mock.Response), &dto.Response); err != nil {
			log.Warning("failed to unmarshal mock response")
		}
	}
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
