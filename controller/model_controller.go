package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	repo "gobackend/repository/database"
	"gobackend/model"
	"gobackend/usecase"
)

//Model type
type Model struct {
	*Controller
	usecase usecase.Model
}

//NewModel New
func NewModel(logger Logger) *Model {
	return &Model{
		Controller: New(logger),
		usecase: usecase.Model{
			Repository: repo.NewModel(),
		},
	}
}

//Create Model
func (controller *Model) Create(c Context) {
  parentID, err := strconv.Atoi(c.Param("parent-id"))
	if err != nil {
		controller.ResponseInvalidRequest(c, err)
		return
	}

	request := model.ModelDTO{}
	if err := c.Bind(&request); err != nil {
		controller.ResponseInvalidRequest(c, err)
		return
	}

	response, err := controller.usecase.Store(parentID, request)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

//Update Model
func (controller *Model) Update(c Context) {
	request := model.ModelDTO{}
	if err := c.Bind(&request); err != nil {
		controller.ResponseWithError(c, err)
		return
	}
  id, err := strconv.Atoi(c.Param("model-id"))
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

//FindAll Model
func (controller *Model) FindAll(c Context) {
	response, err := controller.usecase.FindAll()
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

//FindAllByModel
func (controller *Model) FindAllByParent(c Context) {
  parentID, err := strconv.Atoi(c.Param("parent-id"))
	if err != nil {
		controller.ResponseInvalidRequest(c, err)
		return
  }

	q := model.ModelQuery{}
	if err := c.ShouldBindQuery(&q); err != nil {
		controller.ResponseInvalidRequest(c, err)
	}

	response, err := controller.usecase.FindAllByParent(&q, parentID)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

//FindByID ModelID
func (controller *Model) FindByID(c Context) {
  id, err := strconv.Atoi(c.Param("model-id"))
	if err != nil {
		controller.ResponseInvalidRequest(c, err)
		return
  }

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

//Delete Model
func (controller *Model) Delete(c Context) {
  id, err := strconv.Atoi(c.Param("model-id"))
	if err != nil {
		controller.ResponseInvalidRequest(c, err)
		return
  }

	err = controller.usecase.Delete(id)
	if err != nil {
		controller.ResponseWithError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
