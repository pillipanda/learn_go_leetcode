package climbing_stairs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	var ipt int = 3
	target := 3

	resp := ClimbingStairs(ipt)
	assert.Equal(t, resp, target)
}
