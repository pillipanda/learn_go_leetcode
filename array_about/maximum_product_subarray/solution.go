package maximum_product_subarray

import (
	"math"
)

func MaximumProductSubarray(ipt []int) int {
	max := math.MinInt64
	imax := 1
	imin := 1

	for index := 0; index < len(ipt) ;index++ {
		if ipt[index] < 0 {
			imax, imin = imin, imax
		}

		if imax * ipt[index] > imax {
			imax *= ipt[index]
		}
		if imin * ipt[index] < imin {
			imin *= ipt[index]
		}

		if imax > max {
			max = imax
		}
	}

	return max
}