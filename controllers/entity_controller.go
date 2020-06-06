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
	requestEntity := domain.Entity{}
	err := c.Bind(&requestEntity)

	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	entity, err := controller.usecase.Store(requestEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, entity)
}

//Create Entity
func (controller *EntityController) Update(c Context) {

	validatedEntity := domain.EntityForUpdate{}
	err := c.Bind(&validatedEntity)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	requestEntity := domain.Entity{}
	err = c.Bind(&requestEntity)
	entity, err := controller.usecase.Update(requestEntity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, entity)
}

//GetAll Entity
func (controller *EntityController) GetAll(c Context) {
	entities, err := controller.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, entities)
}

//GetByID EntityID
func (controller *EntityController) GetByID(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Entity, err := controller.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, Entity)
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
