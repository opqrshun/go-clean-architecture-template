package openapi

import (
	"go-clean-architecture/controllers"

	"github.com/gin-gonic/gin"
)

var EntityController = controllers.NewEntityController()

func Create(c *gin.Context) {
	EntityController.Create(c)
}

func GetAll(c *gin.Context) {
	EntityController.GetAll(c)
}

func GetByID(c *gin.Context) {
	EntityController.GetByID(c)
}

// Delete -
func Delete(c *gin.Context) {
	EntityController.Delete(c)
}
