package usecase

import (
	// "fmt"
	"github.com/ttaki/go-clean-architecture-template/model"
)

type Entity struct {
	Repository EntityRepository
}

//Store
func (usecase *Entity) Store(dto model.EntityDTO) (model.Entity, error) {
	p := model.ToEntity(dto)
	id, err := usecase.Repository.Store(p)
	if err != nil {
		return model.Entity{}, err
	}
	//model
	r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

// Update
func (usecase *Entity) Update(dto model.EntityDTO) (model.Entity, error) {
	m, err := usecase.GetTargetEntity(dto.ID)
	if err != nil {
		return model.Entity{}, err
	}
	m.SetRequest(dto)

	id, err := usecase.Repository.Update(&m)
	if err != nil {
		return model.Entity{}, err
	}

	//model
	r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

//FindByID
func (usecase *Entity) FindByID(id int) (model.Entity, error) {
	r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Entity) FindAll(q *model.EntityQuery) ([]model.Entity, error) {
	r, err := usecase.Repository.FindAllFull(q)
	return r, err
}

// GetTargetEntity
func (usecase *Entity) GetTargetEntity(id int) (model.Entity, error) {
	m, err := usecase.Repository.FindByID(id)
	return m, err
}
