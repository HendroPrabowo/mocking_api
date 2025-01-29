package cms

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Mock struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Method       string `gorm:"not null"`
	Path         string `gorm:"unique;not null"`
	ResponseCode int    `gorm:"not null"`
	Request      string `gorm:"type:text"`
	Response     string `gorm:"not null"`
}

func (m *Mock) fromDto(dto MockDto) {
	m.ID = dto.Id
	m.Name = dto.Name
	m.Method = dto.Method
	m.Path = dto.Path
	m.ResponseCode = dto.ResponseCode

	if dto.Response != nil {
		marshal, _ := json.Marshal(dto.Response)
		m.Response = string(marshal)
	}

	if dto.Request != nil {
		marshal, _ := json.Marshal(dto.Request)
		m.Request = string(marshal)
	}
}
