package longest_increasing_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestLongestIncreasingSubsequence(t *testing.T) {
	slice := []int{-2, -1}
	target := 2

	resp := LongestIncreasingSubsequence(slice)
	assert.Equal(t, target, resp)
}
