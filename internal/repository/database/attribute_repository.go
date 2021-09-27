package database

import (
	"gobackend/model"
)

type Attribute struct {
	*Repository
}

//NewAttribute
func NewAttribute() *Attribute {
	return &Attribute{Repository: New()}
}

//FindByID
func (repo *Attribute) FindByID(id int) (model.Attribute, error) {
  var m model.Attribute
  err := repo.FindBaseByID(&m, id)
	return m, err
}

// FindAll
func (repo *Attribute) FindAll() ([]model.Attribute, error) {
  var s []model.Attribute
  err := repo.Find(&s)
	return s, err
}

//FindFullByID
func (repo *Attribute) FindFullByID(id int) (model.Attribute, error) {
  var m model.Attribute
	m.ID = id
  err := repo.FindWithPreload(&m)
	return m, err
}

//FindAll
func (repo *Attribute) FindAllFull() ([]model.Attribute, error) {
  var s []model.Attribute
  err := repo.FindAllWithPreload(s)
	return s, err
}

//FindAllFullByEntity
func (repo *Attribute) FindAllFullByEntity(q *model.AttributeQuery, entityID int) ([]model.Attribute, error) {
  var s []model.Attribute

	condition, params := repo.BuildConditionParams(q)
	condition = condition + "AND attributes.entity_id = ? "
	params = append(params, entityID)
	offset, limit := repo.GetOffsetLimit(q.Page, q.PageSize)

  err := repo.FindAllByConditionWithPreload(&s, condition, params, offset, limit)

	return s, err
}

func (repo *Attribute) BuildConditionParams(q *model.AttributeQuery) (condition string, params []interface{}) {
	// build condition deleted_at=0
	condition = "attributes.deleted_at is NULL "
	if q.Query != "" {
		condition = condition + "AND body LIKE ? "
		params = append(params, "%"+q.Query+"%")
	}
	return
}
