package usecase

import "go-clean-architecture/domain"

type EntityRepository interface {
	Store(domain.Entity) (int, error)
	Update(domain.Entity) (int, error)
	GetById(int) (domain.Entity, error)
	GetAll() ([]domain.Entity, error)
	Delete(int) error
}
