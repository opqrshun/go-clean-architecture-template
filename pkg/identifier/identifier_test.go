package identifier

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	id := "01f5sq3nh10nw9j1zeeg5t17he"
	s := []byte(id)
	var identifier ID
	err := identifier.UnmarshalJSON(s)
	assert.NoError(t, err)
}
