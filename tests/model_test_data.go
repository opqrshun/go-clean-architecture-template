package tests

import "time"

type ModelCreate struct {
	Body string `json:"body" fake:"{sentence:10}`
}

type ModelUpdate struct {
	Body string `json:"body" fake:"{sentence:13}`
}

type Model struct {
	ID int `json:"id,omitempty"`

	Body      string `json:"body" fake:"{sentence:10}`
	ParentID string `json:"parent_id,omitempty"`

	UpdatedAt time.Time `json:"updated_time,omitempty"`
	CreatedAt time.Time `json:"created_time,omitempty"`
}

type ModelInvalid struct {
	Text string `json:"text"`
}

// type InvalidModel struct {
// 	ID int `json:"id"`
// }
