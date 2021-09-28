package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	repo "gobackend/internal/repository/database"
	"gobackend/model"
	"gobackend/internal/usecase"
)

//Attribute type
type Attribute struct {
	*Controller
	usecase usecase.Attribute
}

//NewAttribute New
func NewAttribute(logger Logger) *Attribute {
	return &Attribute{
		Controller: New(logger),
		usecase: usecase.Attribute{
			Repository: repo.NewAttribute(),
		},
	}
}

//Create Attribute
func (controller *Attribute) Create(c Context) {
  entityID, err := strconv.Atoi(c.Param("entity-id"))
	if err != nil {
		controller.RespondInvalidRequest(c, err)
		return
	}

	dto := model.AttributeDTO{}
	if err := c.Bind(&dto); err != nil {
		controller.RespondInvalidRequest(c, err)
		return
	}
  dto.EntityID = entityID

	response, err := controller.usecase.Store(dto)
	if err != nil {
		controller.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

//Update Attribute
func (controller *Attribute) Update(c Context) {
	dto := model.AttributeDTO{}
	if err := c.Bind(&dto); err != nil {
		controller.RespondWithError(c, err)
		return
	}
  id, err := strconv.Atoi(c.Param("attribute-id"))
	if err != nil {
		controller.RespondInvalidRequest(c, err)
		return
  }
	dto.ID = id


	response, err := controller.usecase.Update(dto)
	if err != nil {
		controller.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

//FindAll Attribute
func (controller *Attribute) FindAll(c Context) {
	response, err := controller.usecase.FindAll()
	if err != nil {
		controller.RespondWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

//FindAllByAttribute
func (controller *Attribute) FindAllByEntity(c Context) {
  entityID, err := strconv.Atoi(c.Param("entity-id"))
	if err != nil {
		controller.RespondInvalidRequest(c, err)
		return
  }

	q := model.AttributeQuery{}
	if err := c.ShouldBindQuery(&q); err != nil {
		controller.RespondInvalidRequest(c, err)
	}

	response, err := controller.usecase.FindAllByEntity(&q, entityID)
	if err != nil {
		controller.RespondWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

//FindByID AttributeID
func (controller *Attribute) FindByID(c Context) {
  id, err := strconv.Atoi(c.Param("attribute-id"))
	if err != nil {
		controller.RespondInvalidRequest(c, err)
		return
  }

	if err != nil {
		controller.RespondInvalidRequest(c, err)
		return
	}

	response, err := controller.usecase.FindByID(id)
	if err != nil {
		controller.RespondWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

//Delete Attribute
func (controller *Attribute) Delete(c Context) {
  id, err := strconv.Atoi(c.Param("attribute-id"))
	if err != nil {
		controller.RespondInvalidRequest(c, err)
		return
  }

	err = controller.usecase.Delete(id)
	if err != nil {
		controller.RespondWithError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
