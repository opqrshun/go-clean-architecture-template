package conv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceHeadString(t *testing.T) {
	s, ok := SliceHeadString("https://cdn.sample.com/key/path", "https://cdn.sample.com"+"/")
	assert.True(t, ok)
	assert.Equal(t, "key/path", s)

	_, ok = SliceHeadString("", "https://cdn.sample.com"+"/")
	assert.False(t, ok)

	_, ok = SliceHeadString("aaa", "https://cdn.sample.com"+"/")
	assert.False(t, ok)
}
