package errors

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
  err := Errorf("not found").NotFound()
  ok := IsNotFound(err)
	assert.True(t, ok)
}
