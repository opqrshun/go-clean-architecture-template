package model

import (
	"gorm.io/gorm"

)

type Model struct {
  Base
	Body string `json:"body,omitempty"`
	ParentID int `json:"parent_id,omitempty"`

	DeletedAt gorm.DeletedAt
}


func (model *Model) SetRequest(dto ModelDTO, parentID int) {
	model.Body = dto.Body
	model.ParentID = parentID
	return
}

func (model *Model) SetUpdateRequest(dto ModelDTO) {
	model.Body = dto.Body
	return
}



type ModelDTO struct {
	ID int

	Body      string `json:"body,omitempty" binding:"required"`
	ParentID int
}

// Parent - A single parent.
type ModelQuery struct {
	Query    string `form:"query"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}
