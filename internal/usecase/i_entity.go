package usecase

import (
	repo "github.com/ttaki/go-clean-architecture-sample/internal/repository/database"
	"github.com/ttaki/go-clean-architecture-sample/model"
)

type EntityRepository interface {
	Store(repo.Base) (int, error)
	Update(repo.Base) (int, error)
	FindByID(int) (model.Entity, error)
	Delete(repo.Base) error
	FindFullByID(int) (model.Entity, error)
	FindAll() ([]model.Entity, error)
	FindAllFull(*model.EntityQuery) ([]model.Entity, error)
}
