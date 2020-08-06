package maximum_subarray

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBestTimeToBuyAndSell(t *testing.T)  {
	ipt := []int{-2,1,-3,4,-1,2,1,-5,4}
	target := 6

	resp := MaxmiumSubArray(ipt)
	assert.Equal(t, target, resp)
}