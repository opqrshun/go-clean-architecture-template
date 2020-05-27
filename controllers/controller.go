package controllers

import (
	"go-clean-architecture/domain"
	repo "go-clean-architecture/infrastructure/file"
	"go-clean-architecture/usecase"
	"strconv"
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
	e := domain.Entity{}
	c.Bind(&e)
	entity, err := controller.usecase.Store(e)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, entity)
}

func (controller *EntityController) GetAll(c Context) {
	entities, err := controller.usecase.GetAll()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, entities)
}

func (controller *EntityController) GetById(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Entity, err := controller.usecase.GetById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Entity)
}
