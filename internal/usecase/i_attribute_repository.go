package usecase

import (
	repo "gobackend/internal/repository/database"
	"gobackend/model"
)

type AttributeRepository interface {
	Store(repo.Base) (int, error)
	Update(repo.Base) (int, error)
	FindByID(int) (model.Attribute, error)
	FindAll() ([]model.Attribute, error)
	FindAllFullByEntity(*model.AttributeQuery, int) ([]model.Attribute, error)
	Delete(repo.Base) error
	FindFullByID(int) (model.Attribute, error)
	FindAllFull() ([]model.Attribute, error)
}
