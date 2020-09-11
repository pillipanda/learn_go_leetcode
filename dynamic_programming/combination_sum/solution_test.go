package combination_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	slice := []int{1, 2, 3}
	target := 4

	resp := CombinationSum(slice, target)
	assert.Equal(t, target, resp)
}
