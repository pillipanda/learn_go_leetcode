package product_of_array_except_self

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	ipt := []int{1,2,3,4}
	target := []int{24,12,8,6}

	resp := productOfArrayExceptSelf(ipt)
	assert.Equal(t, target, resp)
}