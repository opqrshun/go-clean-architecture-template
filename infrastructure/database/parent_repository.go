package database

import (
	// "fmt"
	"gobackend/model"
)

type Parent struct {
	*Repository
}

func NewParent() *Parent {
	return &Parent{Repository: New()}
}

//FindByID
func (repo *Parent) FindByID(id int) (model.Parent, error) {
  var c model.Parent
  err := repo.FindBaseByID(&c, id)
	return c, err
}

//get parent
func (repo *Parent) FindFullByID(id int) (model.Parent, error) {
  var parent model.Parent
	parent.ID = id
  err := repo.FindWithPreload(&parent)
	return parent, err
}

//search parents
func (repo *Parent) FindAllFull(q *model.ParentQuery) ([]model.Parent, error) {
  var parents []model.Parent
	condition, params := repo.BuildConditionParams(q)
	offset, limit := repo.GetOffsetLimit(q.Page, q.PageSize)

  err := repo.FindAllByConditionWithPreload(&parents, condition, params, offset, limit)
	return parents, err
}


//BuildConditionParams
func (repo *Parent) BuildConditionParams(q *model.ParentQuery) (condition string, params []interface{}) {
	//TODO seaach body like
	// build condition deleted_at=0
	condition = "parents.deleted_at is NULL "

	if q.Query != "" {
		condition = condition + "AND ( body LIKE ? ) "
		params = append(params, "%"+q.Query+"%")
		params = append(params, "%"+q.Query+"%")
	}

	return
}


//FindAll
func (repo *Parent) FindAll() ([]model.Parent, error) {
  var parents []model.Parent
  err := repo.Find(&parents)
	return parents, err
}

