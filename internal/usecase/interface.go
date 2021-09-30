package usecase

import (
	"github.com/ttaki/go-clean-architecture-template/model"
)

type ExtAPIHandler interface {
	Verify(model.Credential) error
}
