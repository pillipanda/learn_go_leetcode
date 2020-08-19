package three_sum

import (
	"fmt"
	"testing"
)

func TestThreeSumNormal(t *testing.T)  {
	ipt := []int{-1, 0, 1, 2, -1, -4}
	respBs := ThreeSumBS(ipt)
	fmt.Println(respBs)
}