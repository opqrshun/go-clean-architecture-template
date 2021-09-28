package database

import ()

type Base interface {
	GetID() int
	SetCurrentTimeUpdatedAt()
	SetID(int)
	// FindAll() ([]domain.User, error)
	// Delete(int) error
}
