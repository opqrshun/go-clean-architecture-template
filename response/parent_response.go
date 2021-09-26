package response

import (
)

type Parent struct {
	Response

	Body            string  `json:"body,omitempty"`
	Latitude        float64 `json:"latitude,omitempty"`
	Longitude       float64 `json:"longitude,omitempty"`
	ParentImageURL string  `json:"parent_image_url,omitempty"`
	ModelsCount   int     `json:"models_count,omitempty"`
	LikesCount      int     `json:"likes_count,omitempty"`

	Models []Model `json:"models,omitempty"`
	// Tags       []Tag `json:"tags,omitempty"`
}
