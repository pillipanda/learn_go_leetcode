package jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	slice := []int{2, 0, 0}
	target := true

	resp := CanJump(slice)
	assert.Equal(t, target, resp)
}
