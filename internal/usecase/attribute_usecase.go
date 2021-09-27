package usecase

import (
	"gobackend/model"
)

type Attribute struct {
	Repository AttributeRepository
}

func (usecase *Attribute) Store(entityID int, dto model.AttributeDTO ) (model.Attribute, error) {
	//TODO init at once
	m := model.Attribute{}
	m.SetRequest(dto, entityID)

	id, err := usecase.Repository.Store(&m)
	if err != nil {
		return model.Attribute{}, err
	}

  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Attribute) Update(dto model.AttributeDTO) (model.Attribute, error) {
	m, err := usecase.GetTargetAttribute(dto.ID)
	if err != nil {
		return model.Attribute{}, err
	}

	m.SetUpdateRequest(dto)
	id, err := usecase.Repository.Update(&m)
	if err != nil {
		return model.Attribute{}, err
	}

  r, err := usecase.Repository.FindFullByID(id)
	return r, err

}

func (usecase *Attribute) FindAll() ([]model.Attribute, error) {
  r, err := usecase.Repository.FindAllFull()
	return r, err
}

func (usecase *Attribute) FindAllByEntity(q *model.AttributeQuery, entityID int) ([]model.Attribute, error) {
  r, err := usecase.Repository.FindAllFullByEntity(q, entityID)
	return r, err
}

func (usecase *Attribute) FindByID(id int) (model.Attribute, error) {
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Attribute) Delete(id int) (error) {
  var m model.Attribute
  m.ID = id
  err := usecase.Repository.Delete(&m)
  return err
}

//GetTargetAttribute get Authenticated User's Attribute by user ID
//ユーザーのコメントがない場合ErrNotFound
func (usecase *Attribute) GetTargetAttribute(id int) (model.Attribute, error) {
  m, err := usecase.Repository.FindByID(id)
	return m, err
}

