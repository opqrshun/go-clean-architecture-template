package request

import (
)

// Parent - A single parent.
type Parent struct {
	//TODO Bodyを必須にする。Bodyを検索対象にする。
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
