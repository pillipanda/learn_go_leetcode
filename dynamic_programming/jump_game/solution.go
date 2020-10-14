package jump_game

func CanJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if dp[i] == false || nums[i] <= 0 {
			continue
		}
		j := nums[i]
		if (i + j) >= len(nums)-1 {
			return true
		}
		for ; j > 0; j-- {
			dp[i+j] = true
		}
	}

	return false
}
