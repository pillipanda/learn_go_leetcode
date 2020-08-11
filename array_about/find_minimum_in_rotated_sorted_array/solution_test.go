package find_minimum_in_rotated_sorted_array
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindMinInRotatedSortedArray(t *testing.T)  {
	ipt := []int{2,3,4,5,1}
	target := 1
	resp := FindMinInRotatedSortedArray(ipt)
	assert.Equal(t, target, resp)

	iptS := []int{4,5,6,7,0,1,2}
	targetS := 0
	respS := FindMinInRotatedSortedArray(iptS)
	assert.Equal(t, targetS, respS)
}