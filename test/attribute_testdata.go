package test

import "time"

type AttributeCreate struct {
	Body string `json:"body" fake:"{sentence:10}`
}

type AttributeUpdate struct {
	Body string `json:"body" fake:"{sentence:13}`
}

type Attribute struct {
	ID int `json:"id,omitempty"`

	Body      string `json:"body" fake:"{sentence:10}`
	EntityID string `json:"entity_id,omitempty"`

	UpdatedAt time.Time `json:"updated_time,omitempty"`
	CreatedAt time.Time `json:"created_time,omitempty"`
}

type AttributeInvalid struct {
	Text string `json:"text"`
}

// type InvalidAttribute struct {
// 	ID int `json:"id"`
// }
