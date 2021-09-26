package usecase

import (
	"gobackend/domain"
	repo "gobackend/infrastructure/database"
	"gobackend/request"
	"gobackend/response"
)

type ParentRepository interface {
	Store(repo.Base) (int, error)
	Update(repo.Base) (int, error)
	FindByID(int) (domain.Parent, error)
	Delete(repo.Base) error
	FindFullByID(int) (response.Parent, error)
	FindAll() ([]domain.Parent, error)
	FindAllFull(*request.ParentQuery) ([]response.Parent, error)
}
