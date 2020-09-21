package houserobber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseRobber(t *testing.T) {
	slice := []int{2, 7, 9, 3, 1}
	target := 12

	resp := HouseRobber(slice)
	assert.Equal(t, target, resp)
}
