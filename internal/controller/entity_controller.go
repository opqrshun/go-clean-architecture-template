package controller

import (
	"net/http"
	"strconv"

	repo "gobackend/internal/repository/database"
	"gobackend/model"
	"gobackend/internal/usecase"
)

//Entity type
type Entity struct {
	*Controller
	usecase usecase.Entity
}

//NewEntity New
func NewEntity(logger Logger) *Entity {
	return &Entity{
		Controller: New(logger),
		usecase: usecase.Entity{
			Repository: repo.NewEntity(),
		},
	}
}

//Create Entity
func (controller *Entity) Create(c Context) {
  //validate
	dto := model.EntityDTO{}
	if err := c.Bind(&dto); err != nil {
    controller.ResponseInvalidRequest(c, err)
    return
	}

  //store
	response, err := controller.usecase.Store(dto)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}


//Update Entity
func (controller *Entity) Update(c Context) {
  //validate
	dto := model.EntityDTO{}
	if err := c.Bind(&dto); err != nil {
    controller.ResponseInvalidRequest(c, err)
		return
	}

  id, err := strconv.Atoi(c.Param("entity-id"))
	if err != nil {
		controller.ResponseInvalidRequest(c, err)
		return
  }

	dto.ID = id


	response, err := controller.usecase.Update(dto)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

//FindAll Entity
func (controller *Entity) FindAll(c Context) {

  //validate
	q := model.EntityQuery{}
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


//FindByID EntityID
func (controller *Entity) FindByID(c Context) {
  id, err := strconv.Atoi(c.Param("entity-id"))
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
