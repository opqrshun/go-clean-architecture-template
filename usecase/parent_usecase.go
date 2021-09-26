package usecase

import (
	// "fmt"
	"gobackend/domain"
	"gobackend/request"
	"gobackend/response"
)

type Parent struct {
	Repository ParentRepository
}

//Store
// with Save tags
func (usecase *Parent) Store(req request.Parent) (response.Parent, error) {
	c := domain.Parent{}
	c.SetRequest(req)
	id, err := usecase.Repository.Store(&c)
	if err != nil {
		return response.Parent{}, err
	}
	//response
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

// Update
func (usecase *Parent) Update(req request.Parent) (response.Parent, error) {
	c, err := usecase.GetTargetParent(req.ID)
	if err != nil {
		return response.Parent{}, err
	}
	c.SetRequest(req)

	id, err := usecase.Repository.Update(&c)
	if err != nil {
		return response.Parent{}, err
	}

	//response
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}


//FindByID
func (usecase *Parent) FindByID(id int) (response.Parent, error) {
  r, err := usecase.Repository.FindFullByID(id)
	return r, err
}

func (usecase *Parent) FindAll(q *request.ParentQuery) ([]response.Parent, error) {
  r, err := usecase.Repository.FindAllFull(q)
	return r, err
}

// GetTargetParent
func (usecase *Parent) GetTargetParent(id int) (domain.Parent, error) {
  parent, err := usecase.Repository.FindByID(id)
	if err != nil {
		return domain.Parent{}, err
	}
	return parent, nil
}
