package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIs(t *testing.T) {
	err := Errorf("not found").NotFound()
	ok := IsNotFound(err)
	assert.True(t, ok)
}
