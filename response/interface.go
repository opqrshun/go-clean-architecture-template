package response

import (
	"time"
)

type Base interface {
	GetID() int
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
