package usecase

import (
	"gobackend/domain"
	repo "gobackend/infrastructure/database"
	"gobackend/request"
	"gobackend/response"
)

type ModelRepository interface {
	Store(repo.Base) (int, error)
	Update(repo.Base) (int, error)
	FindByID(int) (domain.Model, error)
	FindAll() ([]domain.Model, error)
	FindAllFullByParent(*request.ModelQuery, int) ([]response.Model, error)
	Delete(repo.Base) error
	FindFullByID(int) (response.Model, error)
	FindAllFull() ([]response.Model, error)
}
