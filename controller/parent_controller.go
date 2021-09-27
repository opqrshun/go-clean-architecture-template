package controller

import (
	"net/http"
	"strconv"

	repo "gobackend/infrastructure/database"
	"gobackend/model"
	"gobackend/usecase"
)

//Parent type
type Parent struct {
	*Controller
	usecase usecase.Parent
}

//NewParent New
func NewParent(logger Logger) *Parent {
	return &Parent{
		Controller: New(logger),
		usecase: usecase.Parent{
			Repository: repo.NewParent(),
		},
	}
}

//Create Parent
func (controller *Parent) Create(c Context) {
  //validate
	request := model.ParentDTO{}
	if err := c.Bind(&request); err != nil {
    controller.ResponseInvalidRequest(c, err)
    return
	}

  //store
	response, err := controller.usecase.Store(request)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}


//Update Parent
func (controller *Parent) Update(c Context) {
  //validate
	request := model.ParentDTO{}
	if err := c.Bind(&request); err != nil {
    controller.ResponseInvalidRequest(c, err)
		return
	}

  id, err := strconv.Atoi(c.Param("parent-id"))
	if err != nil {
		controller.ResponseInvalidRequest(c, err)
		return
  }

	request.ID = id


	response, err := controller.usecase.Update(request)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

//FindAll Parent
func (controller *Parent) FindAll(c Context) {

  //validate
	q := model.ParentQuery{}
	if err := c.ShouldBindQuery(&q); err != nil {
    controller.ResponseInvalidRequest(c, err)
    return
	}

	response, err := controller.usecase.FindAll(&q)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}


//FindByID ParentID
func (controller *Parent) FindByID(c Context) {
  id, err := strconv.Atoi(c.Param("parent-id"))
	if err != nil {
		controller.ResponseInvalidRequest(c, err)
		return
  }

	response, err := controller.usecase.FindByID(id)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
