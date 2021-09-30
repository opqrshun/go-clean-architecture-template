package extapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	h := NewAPIHandler()

	err := h.Ping()
	assert.NoError(t, err, "Should be no error")
}
