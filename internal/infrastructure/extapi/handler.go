package extapi

import (
	// "log"
	"net/http"

	"github.com/ttaki/go-clean-architecture-template/internal/config"
	_service "github.com/ttaki/go-clean-architecture-template/internal/infrastructure/extapi/extapiservice"
	"github.com/ttaki/go-clean-architecture-template/pkg/errors"
)

type APIHandler struct {
	client  *http.Client
	URL     string
	service *_service.Service
}

var client *http.Client

func init() {
	client = newClient()
}

func newClient() *http.Client {
	return &http.Client{}
}

func getClient() *http.Client {
	return client
}

func NewAPIHandler() *APIHandler {
	config := config.GetExtAPIConfig()

	return &APIHandler{
		client:  getClient(),
		URL:     config.URL,
		service: _service.New(getClient()),
	}
}

func (h *APIHandler) Ping() error {
	req, err := http.NewRequest("Ping", h.URL, nil)
	if err != nil {
		return BuildAPIError(err, "Ping")
	}

	_, err = client.Do(req)
	if err != nil {
		return BuildAPIError(err, "Ping")
	}

	return nil
}

func BuildAPIError(err error, methodName string) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, http.ErrNotSupported) {
		return errors.Wrapf(err, "extapi error, method: "+methodName).Internal()
	}
	return errors.Wrapf(err, "extapi error, method: "+methodName).Internal()
}
