package usecase

import (
	repo "gobackend/infrastructure/database"
	"gobackend/model"
)

type ModelRepository interface {
	Store(repo.Base) (int, error)
	Update(repo.Base) (int, error)
	FindByID(int) (model.Model, error)
	FindAll() ([]model.Model, error)
	FindAllFullByParent(*model.ModelQuery, int) ([]model.Model, error)
	Delete(repo.Base) error
	FindFullByID(int) (model.Model, error)
	FindAllFull() ([]model.Model, error)
}
