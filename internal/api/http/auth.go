package http

import (
	"github.com/gin-gonic/gin"
	"gobackend/internal/controller"
	"gobackend/pkg/errors"
	"gobackend/pkg/logger"
)

var Controller = controller.New(logger.GetLogger())

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

    accessKey := ""
    var err error

		if err != nil {
			aerr := errors.Wrapf(err, "token is invalid, method: Auth").Unauthorized()
			Controller.ResponseWithError(c, aerr)
			c.Abort()
			return
		}

		c.Set("accessKey", accessKey)
		c.Next()

	}
}


type Header struct {
	accessKey string `header:"authorization" binding:"required"`
}

