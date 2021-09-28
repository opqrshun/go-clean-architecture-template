package test

import "time"

type EntityCreate struct {
	Body string `json:"body" fake:"{sentence:3}"`
}

type EntityUpdate struct {
	ID   int
	Body string `json:"body" fake:"{sentence:2}"`
}

type Entity struct {
	ID int `json:"id,omitempty"`

	Body       string      `json:"body" fake:"{sentence:3}"`
	Attributes []Attribute `json:"attributes,omitempty fake:"skip"`

	CreatedAt time.Time `json:"created_time,omitempty"`
	UpdatedAt time.Time `json:"updated_time,omitempty"`
}

type EntityInvalid struct {
	ID string `json:"id"`
}

// type InvalidEntity struct {
// 	ID string `json:"id"`
// }
