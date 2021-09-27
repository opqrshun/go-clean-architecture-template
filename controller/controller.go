package controller

import (
	"net/http"

	"gobackend/core/errors"
)

type H map[string]interface{}

//Parent type
type Controller struct {
	Logger Logger
}

//NewParent New
func New(logger Logger) *Controller {
	return &Controller{
		Logger: logger,
	}
}

// ResponseWithError ...
func (controller *Controller) ResponseWithError(c Context, err error) {
	var nextMsg string
	if next := errors.Unwrap(err); next != nil {
		nextMsg = next.Error()
	}

	var aerr *errors.AppError
	if errors.As(err, &aerr) {
		c.JSON(aerr.HttpStatus(), H{
			"code":    aerr.Code(),
			"message": aerr.DisplayMessage(),
		})
		controller.Logger.Errorw("method: ResponseWithError", "code", aerr.Code(), "msg", aerr.Error(), "childErr", nextMsg)
	} else {
		c.JSON(http.StatusInternalServerError, H{
			"code":    "unknown",
			"message": "unknown",
		})
	}
	return
}

func (controller *Controller) ResponseInvalidRequest(c Context,err error) {
  controller.ResponseWithError(c, errors.Errorf("invalid request, err: %v", err).InvalidRequest())
}

//GetAuthorizedID
func GetAuthorizedID(c Context) string {
	return c.MustGet("idpid").(string)
}
