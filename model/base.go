package model

import (
	"time"
)

type Base struct {
	ID        int       `json:"id,omitempty" gorm:"primary_key;"`
	CreatedAt time.Time `json:"created_time,omitempty"`
	UpdatedAt time.Time `json:"updated_time,omitempty"`
}

func (base *Base) GetID() int {
	return base.ID
}

func (base *Base) SetID(id int) {
	base.ID = id
}

func (base *Base) GetCreatedAt() time.Time {
	return base.CreatedAt
}

func (base *Base) GetUpdatedAt() time.Time {
	return base.UpdatedAt
}

func (base *Base) SetCurrentTimeUpdatedAt() {
	base.UpdatedAt = time.Now()
}
