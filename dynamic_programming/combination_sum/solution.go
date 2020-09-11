package combination_sum

func CombinationSum(nums []int, target int) int {
	dp := make([]int, target+1)
	valid_nums := []int{}
	for _, num := range nums {
		if num <= target {
			valid_nums = append(valid_nums, num)
			dp[num] = 1
		}
	}

	for i := 0; i <= target; i++ {
		for _, num := range valid_nums {
			if (i - num) > 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
