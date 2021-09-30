package extapiservice

import (
	"net/http"

	apimodel "github.com/ttaki/go-clean-architecture-template/internal/infrastructure/extapi/extapimodel"
)

// Entity
// Actual implementation of API and DB access
// Hide complex processing here.
// The facade pattern is a possible example.
type Service struct {
	client *http.Client
}

func New(c *http.Client) *Service {
	return &Service{
		client: c,
	}
}

func (s *Service) Complex(m *apimodel.User) (apimodel.User, error) {
	return apimodel.User{}, nil
}
