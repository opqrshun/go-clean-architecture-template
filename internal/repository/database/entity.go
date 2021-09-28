package database

import (
	// "fmt"
	"github.com/ttaki/go-clean-architecture-template/model"
)

type Entity struct {
	*Repository
}

func NewEntity() *Entity {
	return &Entity{Repository: New()}
}

//FindByID
func (repo *Entity) FindByID(id int) (model.Entity, error) {
	var m model.Entity
	err := repo.FindBaseByID(&m, id)
	return m, err
}

//FindFullByID
func (repo *Entity) FindFullByID(id int) (model.Entity, error) {
	var m model.Entity
	m.ID = id
	err := repo.FindWithPreload(&m)
	return m, err
}

//FindAllFull
func (repo *Entity) FindAllFull(q *model.EntityQuery) ([]model.Entity, error) {
	var s []model.Entity
	condition, params := repo.BuildConditionParams(q)
	offset, limit := repo.GetOffsetLimit(q.Page, q.PageSize)

	err := repo.FindAllByConditionWithPreload(&s, condition, params, offset, limit)
	return s, err
}

//BuildConditionParams
func (repo *Entity) BuildConditionParams(q *model.EntityQuery) (condition string, params []interface{}) {
	//TODO seaach body like
	// build condition deleted_at=0
	condition = "entities.deleted_at is NULL "

	if q.Query != "" {
		condition = condition + "AND ( body LIKE ? ) "
		params = append(params, "%"+q.Query+"%")
		params = append(params, "%"+q.Query+"%")
	}

	return
}

//FindAll
func (repo *Entity) FindAll() ([]model.Entity, error) {
	var s []model.Entity
	err := repo.Find(&s)
	return s, err
}
