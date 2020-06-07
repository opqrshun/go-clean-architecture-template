package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-clean-architecture/domain"
	repo "go-clean-architecture/infrastructure/memory"
	"go-clean-architecture/usecase"
)

//EntityController type
type EntityController struct {
	usecase usecase.EntityUsecase
}

//NewEntityController New
func NewEntityController() *EntityController {
	return &EntityController{
		usecase: usecase.EntityUsecase{
			EntityRepository: repo.NewEntityRepository(),
		},
	}
}

//Create Entity
func (controller *EntityController) Create(c Context) {
	request := domain.Entity{}
	err := c.Bind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	response, err := controller.usecase.Store(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

//Update Entity
func (controller *EntityController) Update(c Context) {

	validated := domain.EntityForUpdate{}
	err := c.Bind(&validated)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}

	request := domain.Entity{}
	err = c.Bind(&request)
	response, err := controller.usecase.Update(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

//GetAll Entity
func (controller *EntityController) GetAll(c Context) {
	response, err := controller.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

//GetById EntityId
func (controller *EntityController) GetById(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := controller.usecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

//Delete Entity
func (controller *EntityController) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.usecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
