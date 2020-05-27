package http

import (
	"go-clean-architecture/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	EntityController := controllers.NewEntityController()

	router.POST("/entities", func(c *gin.Context) { EntityController.Create(c) })
	router.GET("/entities", func(c *gin.Context) { EntityController.GetAll(c) })
	router.GET("/entities/:id", func(c *gin.Context) { EntityController.GetById(c) })

	Router = router
}
