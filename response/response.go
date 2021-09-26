package response

import (
	"time"

)

type Response struct {
	// gorm.Model
	ID        int `json:"id,omitempty"`
	CreatedAt time.Time     `json:"created_time,omitempty"`
	UpdatedAt time.Time     `json:"updated_time,omitempty"`
}

func (response *Response) SetModel(base Base) {
	response.ID = base.GetID()
	response.CreatedAt = base.GetCreatedAt()
	response.UpdatedAt = base.GetUpdatedAt()
}
