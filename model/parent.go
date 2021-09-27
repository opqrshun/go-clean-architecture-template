package model
import (
	"gorm.io/gorm"

)

// Parent - A single parent.
type Parent struct {
	Base
	Body            string  `json:"body,omitempty"`

	DeletedAt gorm.DeletedAt
	Models []Model `json:"models,omitempty"`
}

func (parent *Parent) SetRequest(dto ParentDTO) {
	parent.Body = dto.Body
}

// Parent - A single parent.
type ParentDTO struct {
	ID    int
	Body  string `json:"body,omitempty"`

	// ParentImageURL string `json:"parent_image_url,omitempty"`
}

// Parent - A single parent.
type ParentQuery struct {
	Query string `form:"query"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}
