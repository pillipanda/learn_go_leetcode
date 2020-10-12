package unique_paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	resp := UniquePaths(3, 7)
	assert.Equal(t, resp, 28)
}
