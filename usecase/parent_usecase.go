package usecase

import (
	// "fmt"
	"gobackend/model"
)

type Parent struct {
	Repository ParentRepository
}

//Store
// with Save tags
func (usecase *Parent) Store(req model.ParentDTO) (model.Parent, error) {
	c := model.Parent{}
	c.SetRequest(req)
	id, err := usecase.Repository.Store(&c)
	if err != nil {
		return model.Parent{}, err
	}
	//model
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

// Update
func (usecase *Parent) Update(req model.ParentDTO) (model.Parent, error) {
	c, err := usecase.GetTargetParent(req.ID)
	if err != nil {
		return model.Parent{}, err
	}
	c.SetRequest(req)

	id, err := usecase.Repository.Update(&c)
	if err != nil {
		return model.Parent{}, err
	}

	//model
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}


//FindByID
func (usecase *Parent) FindByID(id int) (model.Parent, error) {
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Parent) FindAll(q *model.ParentQuery) ([]model.Parent, error) {
  r, err := usecase.Repository.FindAllFull(q)
	return r, err
}

// GetTargetParent
func (usecase *Parent) GetTargetParent(id int) (model.Parent, error) {
  parent, err := usecase.Repository.FindByID(id)
	if err != nil {
		return model.Parent{}, err
	}
	return parent, nil
}
