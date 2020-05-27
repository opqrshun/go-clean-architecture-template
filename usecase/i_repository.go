package usecase

import "go-clean-architecture/domain"

type EntityRepository interface {
	Store(domain.Entity) (int, error)
	FindById(int) (domain.Entity, error)
	FindAll() ([]domain.Entity, error)
}
