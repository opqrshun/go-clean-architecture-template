package usecase

import (
	"gobackend/model"
)

type Model struct {
	Repository ModelRepository
}

func (usecase *Model) Store(parentID int, req model.ModelDTO ) (model.Model, error) {
	//TODO init at once
	m := model.Model{}
	m.SetRequest(req, parentID)

	id, err := usecase.Repository.Store(&m)
	if err != nil {
		return model.Model{}, err
	}

  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Model) Update(req model.ModelDTO) (model.Model, error) {
	m, err := usecase.GetTargetModel(req.ID)
	if err != nil {
		return model.Model{}, err
	}

	m.SetUpdateRequest(req)
	id, err := usecase.Repository.Update(&m)
	if err != nil {
		return model.Model{}, err
	}

  r, err := usecase.Repository.FindFullByID(id)
	return r, err

}

func (usecase *Model) FindAll() ([]model.Model, error) {
  r, err := usecase.Repository.FindAllFull()
	return r, err
}

func (usecase *Model) FindAllByParent(q *model.ModelQuery, parentID int) ([]model.Model, error) {
  r, err := usecase.Repository.FindAllFullByParent(q, parentID)
	return r, err
}

func (usecase *Model) FindByID(id int) (model.Model, error) {
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Model) Delete(id int) (error) {
  var model model.Model
  model.ID = id
  err := usecase.Repository.Delete(&model)
  return err
}

//GetTargetModel get Authenticated User's Model by user ID
//ユーザーのコメントがない場合ErrNotFound
func (usecase *Model) GetTargetModel(id int) (model.Model, error) {
  m, err := usecase.Repository.FindByID(id)
	return m, err
}

