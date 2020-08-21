package container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {
	ipt := []int{1,8,6,2,5,4,8,3,7}
	target := 49

	resp := ContainerWithMostWater(ipt)
	assert.Equal(t, target, resp)
}