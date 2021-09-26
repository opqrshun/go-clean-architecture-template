package usecase

import (
	"gobackend/domain"
	"gobackend/request"
	"gobackend/response"
)

type Model struct {
	Repository ModelRepository
}

func (usecase *Model) Store(parentID int, req request.Model ) (response.Model, error) {
	//TODO init at once
	model := domain.Model{}
	model.SetRequest(req, parentID)

	id, err := usecase.Repository.Store(&model)
	if err != nil {
		return response.Model{}, err
	}

  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Model) Update(req request.Model) (response.Model, error) {
	model, err := usecase.GetTargetModel(req.ID)
	if err != nil {
		return response.Model{}, err
	}

	model.SetUpdateRequest(req)

	id, err := usecase.Repository.Update(&model)
	if err != nil {
		return response.Model{}, err
	}

  r, err := usecase.Repository.FindFullByID(id)
	return r, err

}

func (usecase *Model) FindAll() ([]response.Model, error) {
  r, err := usecase.Repository.FindAllFull()
	return r, err
}

func (usecase *Model) FindAllByParent(q *request.ModelQuery, parentID int) ([]response.Model, error) {
  r, err := usecase.Repository.FindAllFullByParent(q, parentID)
	return r, err
}

func (usecase *Model) FindByID(id int) (response.Model, error) {
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Model) Delete(id int) (error) {
  var model domain.Model
  model.ID = id
  err := usecase.Repository.Delete(&model)
  return err
}

//GetTargetModel get Authenticated User's Model by user ID
//ユーザーのコメントがない場合ErrNotFound
func (usecase *Model) GetTargetModel(id int) (domain.Model, error) {
  model, err := usecase.Repository.FindByID(id)
	if err != nil {
		return domain.Model{} , err
	}

	return model, nil
}

