package tests

import "time"

type ParentCreate struct {
	Body     string   `json:"body" fake:"{sentence:3}"`
}

type ParentUpdate struct {
	ID       int 
	Body     string   `json:"body" fake:"{sentence:2}"`
}

type Parent struct {
	ID int `json:"id,omitempty"`


	Body           string `json:"body" fake:"{sentence:3}"`
  Models []Model `json:"models,omitempty fake:"skip"`

	CreatedAt time.Time `json:"created_time,omitempty"`
	UpdatedAt time.Time `json:"updated_time,omitempty"`
}

type ParentInvalid struct {
	ID string `json:"id"`
}

// type InvalidParent struct {
// 	ID string `json:"id"`
// }
