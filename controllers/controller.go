package controllers

import (
	"go-clean-architecture/domain"
	repo "go-clean-architecture/infrastructure/memory"
	"go-clean-architecture/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntityController struct {
	usecase usecase.Usecase
}

func NewEntityController() *EntityController {
	return &EntityController{
		usecase: usecase.Usecase{
			EntityRepository: repo.NewEntityRepository(),
		},
	}
}

func (controller *EntityController) Create(c Context) {
	request_entity := domain.Entity{}
	c.Bind(&request_entity)
	entity, err := controller.usecase.Store(request_entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, entity)
}

func (controller *EntityController) GetAll(c Context) {
	entities, err := controller.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, entities)
}

func (controller *EntityController) GetById(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Entity, err := controller.usecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, Entity)
}

func (controller *EntityController) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.usecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
