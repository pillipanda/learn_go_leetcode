package two_sum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{1, 0}

	resp := TwoSum(nums, target)
	assert.Equal(t, resp, expected)
}
