package controller

import (
	"net/http"

	"gobackend/pkg/errors"
)

type H map[string]interface{}

//Entity type
type Controller struct {
	Logger Logger
}

//NewEntity New
func New(logger Logger) *Controller {
	return &Controller{
		Logger: logger,
	}
}

// RespondWithError ...
func (controller *Controller) RespondWithError(c Context, err error) {
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
		controller.Logger.Errorw("method: RespondWithError", "code", aerr.Code(), "msg", aerr.Error(), "childErr", nextMsg)
	} else {
		c.JSON(http.StatusInternalServerError, H{
			"code":    "unknown",
			"message": "unknown",
		})
	}
	return
}

func (controller *Controller) RespondInvalidRequest(c Context,err error) {
  controller.RespondWithError(c, errors.Errorf("invalid request, err: %v", err).InvalidRequest())
}

//GetAuthorizedID
func GetAuthorizedID(c Context) string {
	return c.MustGet("idpid").(string)
}
