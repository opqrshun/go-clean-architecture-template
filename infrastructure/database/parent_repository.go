package database

import (
	// "fmt"
	"gobackend/domain"
	"gobackend/request"
	"gobackend/response"
)

type Parent struct {
	*Repository
}

func NewParent() *Parent {
	return &Parent{Repository: New()}
}

//FindByID
func (repo *Parent) FindByID(id int) (domain.Parent, error) {
  var c domain.Parent
  err := repo.FindBaseByID(&c, id)
	return c, err
}

//get parent
func (repo *Parent) FindFullByID(id int) (response.Parent, error) {
  var parent response.Parent
	parent.ID = id
  err := repo.FindWithPreload(&parent)
	return parent, err
}

//search parents
func (repo *Parent) FindAllFull(q *request.ParentQuery) ([]response.Parent, error) {
  var parents []response.Parent
	condition, params := repo.BuildConditionParams(q)
	offset, limit := repo.GetOffsetLimit(q.Page, q.PageSize)

  err := repo.FindAllByConditionWithPreload(&parents, condition, params, offset, limit)
	return parents, err
}


//BuildConditionParams
func (repo *Parent) BuildConditionParams(q *request.ParentQuery) (condition string, params []interface{}) {
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
func (repo *Parent) FindAll() ([]domain.Parent, error) {
  var parents []domain.Parent
  err := repo.Find(&parents)
	return parents, err
}

