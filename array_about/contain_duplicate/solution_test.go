package contain_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainDuplicate(t *testing.T) {
	ipt := []int{1, 2, 3, 1}
	target := true

	resp := ContainDuplicate(ipt)
	assert.Equal(t, resp, target)
}