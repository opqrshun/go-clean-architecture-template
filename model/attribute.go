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

//ToAttribute convert to Attribute, return Attribute pointer
func ToAttribute(dto AttributeDTO) *Attribute{
  return &Attribute{
    EntityID:      dto.EntityID,
		Body:          dto.Body,
  }
}


func (attribute *Attribute) SetDTO(dto AttributeDTO) {
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
