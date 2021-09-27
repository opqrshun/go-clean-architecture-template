package model

import (
	"gorm.io/gorm"

)

// Entity - A single entity.
type Entity struct {
	Base
	Body            string  `json:"body,omitempty"`

	DeletedAt gorm.DeletedAt
	Attributes []Attribute `json:"attributes,omitempty"`
}

func (entity *Entity) SetRequest(dto EntityDTO) {
	entity.Body = dto.Body
}

// Entity - A single entity.
type EntityDTO struct {
	ID    int
	Body  string `json:"body,omitempty"`

	// EntityImageURL string `json:"entity_image_url,omitempty"`
}

// Entity - A single entity.
type EntityQuery struct {
	Query string `form:"query"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}
