package model

import (
	"gorm.io/gorm"

)

type Attribute struct {
  Base
	Body string `json:"body,omitempty"`
	EntityID int `json:"entity_id,omitempty"`

	DeletedAt gorm.DeletedAt
}


func (attribute *Attribute) SetRequest(dto AttributeDTO, entityID int) {
	attribute.Body = dto.Body
	attribute.EntityID = entityID
	return
}

func (attribute *Attribute) SetUpdateRequest(dto AttributeDTO) {
	attribute.Body = dto.Body
	return
}



type AttributeDTO struct {
	ID int

	Body      string `json:"body,omitempty" binding:"required"`
	EntityID int
}

// Entity - A single entity.
type AttributeQuery struct {
	Query    string `form:"query"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}
