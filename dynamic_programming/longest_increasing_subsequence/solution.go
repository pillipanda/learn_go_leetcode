package longest_increasing_subsequence

import (
	"math"
)

func LongestIncreasingSubsequence (ipt_nums []int) int {
	if len(ipt_nums) < 1 {
		return 0
	}

	max := 1
	by_now_min_val := math.MaxInt32
	index_icreaseamount := make([]int, len(ipt_nums))
	for i := 0; i < len(ipt_nums); i ++ {
		val := ipt_nums[i]
		index_icreaseamount[i] = math.MinInt32
		if val <= by_now_min_val {
			index_icreaseamount[i] = 1
			by_now_min_val = val
			continue
		}
		for j := 0; j < i; j ++ {
			if ipt_nums[j] >= val {
				continue
			}
			// 由子问题的求解过程
			if index_icreaseamount[j] + 1 > index_icreaseamount[i] {
				index_icreaseamount[i] = index_icreaseamount[j] + 1
			}
		}
		if index_icreaseamount[i] > max {
			max = index_icreaseamount[i]
		}
	}
	return max
}