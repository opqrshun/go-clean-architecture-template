package extapi

import (
	// "log"
	"net/http"

	"github.com/ttaki/go-clean-architecture-template/internal/config"
	"github.com/ttaki/go-clean-architecture-template/pkg/errors"
)

type APIHandler struct {
	client *http.Client
	URL    string
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
		client: getClient(),
		URL:    config.URL,
	}
}

func (h *APIHandler) Find() error {
	req, err := http.NewRequest("GET", h.URL, nil)
	if err != nil {
		return BuildAPIError(err, "Find")
	}

	_, err = client.Do(req)
	if err != nil {
		return BuildAPIError(err, "Find")
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
