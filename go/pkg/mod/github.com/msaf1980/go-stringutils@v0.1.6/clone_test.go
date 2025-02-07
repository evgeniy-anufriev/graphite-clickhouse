package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CloneString(t *testing.T) {
	in := "Hello, World!"
	out := Clone(in)
	assert.Equal(t, "Hello, World!", out)
}

func Test_CloneBytes(t *testing.T) {
	in := []byte("Hello, World!")
	out := CloneBytes(in)
	assert.Equal(t, "Hello, World!", string(out))
}
