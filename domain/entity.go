package domain

import "time"

type Entity struct {
	Id        int
	DeletedAt time.Time
}
