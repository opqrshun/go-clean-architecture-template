package response

import (
	"gobackend/domain"
)

type Model struct {
	Response
	Body      string        `json:"body,omitempty"`
	ParentID int `json:"parent_id,omitempty"`
}

func (res *Model) SetModel(c *domain.Model) {
	res.SetModel(c)

	res.Body = c.Body
	res.ParentID = c.ParentID
}
