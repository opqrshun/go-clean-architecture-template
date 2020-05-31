package domain

import "time"

type Entity struct {
	Id   int    `json:"id"`
	Text string `json:"text"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
