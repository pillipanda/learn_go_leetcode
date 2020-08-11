package maximum_product_subarray


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaximumProductSubarray(t *testing.T)  {
	ipt := []int{2,3,-2,4,3}
	target := 12

	resp := MaximumProductSubarray(ipt)
	assert.Equal(t, target, resp)
}