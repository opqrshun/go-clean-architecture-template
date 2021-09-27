/*
 * Swagger MapSNS
 *
 * gobackend api
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://http-generator.tech)
 */

package http

import (
	"gobackend/internal/controller"
	"github.com/gin-gonic/gin"
	"gobackend/pkg/logger"
)

var Model = controller.NewModel(logger.GetLogger())

// CreateModel -
func CreateModel(c *gin.Context) {
	Model.Create(c)
}

// DeleteModel -
func DeleteModel(c *gin.Context) {
	Model.Delete(c)
}

// GetModel -
func GetModel(c *gin.Context) {
	Model.FindByID(c)
}

// GetModelsByParent - Your GET endpoint
func GetModelsByParent(c *gin.Context) {
	Model.FindAllByParent(c)
}

// UpdateModel -
func UpdateModel(c *gin.Context) {
	Model.Update(c)
}