package usecase

import (
	repo "gobackend/internal/repository/database"
	"gobackend/model"
)

type ParentRepository interface {
	Store(repo.Base) (int, error)
	Update(repo.Base) (int, error)
	FindByID(int) (model.Parent, error)
	Delete(repo.Base) error
	FindFullByID(int) (model.Parent, error)
	FindAll() ([]model.Parent, error)
	FindAllFull(*model.ParentQuery) ([]model.Parent, error)
}
