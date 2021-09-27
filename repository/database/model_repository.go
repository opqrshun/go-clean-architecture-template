package database

import (
	"gobackend/model"
)

type Model struct {
	*Repository
}

func NewModel() *Model {
	return &Model{Repository: New()}
}

func (repo *Model) FindByID(id int) (model.Model, error) {
  var model model.Model
  err := repo.FindBaseByID(&model, id)
	return model, err
}

func (repo *Model) FindAll() ([]model.Model, error) {
  var models []model.Model
  err := repo.Find(&models)
	return models, err
}

func (repo *Model) FindFullByID(id int) (model.Model, error) {
  var model model.Model
	model.ID = id
  err := repo.FindWithPreload(&model)
	return model, err
}

//FindAll
func (repo *Model) FindAllFull() ([]model.Model, error) {
  var models []model.Model
  err := repo.FindAllWithPreload(models)
	return models, err
}

//search parents by Authenticated user or other user
func (repo *Model) FindAllFullByParent(q *model.ModelQuery, parentID int) ([]model.Model, error) {
  var models []model.Model

	condition, params := repo.BuildConditionParams(q)
	condition = condition + "AND models.parent_id = ? "
	params = append(params, parentID)
	offset, limit := repo.GetOffsetLimit(q.Page, q.PageSize)

  err := repo.FindAllByConditionWithPreload(&models, condition, params, offset, limit)

	return models, err
}

func (repo *Model) BuildConditionParams(q *model.ModelQuery) (condition string, params []interface{}) {
	// build condition deleted_at=0
	condition = "models.deleted_at is NULL "
	if q.Query != "" {
		condition = condition + "AND body LIKE ? "
		params = append(params, "%"+q.Query+"%")
	}
	return
}
