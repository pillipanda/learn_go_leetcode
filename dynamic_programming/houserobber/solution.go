package houserobber

func HouseRobber(nums []int) int {
	dp := make([]int, len(nums))
	max := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
			max = dp[i]
			continue
		}
		if i == 1 {
			if nums[i] >= nums[i-1] {
				dp[i] = nums[i]
				max = dp[i]
			} else {
				dp[i] = dp[i-1]
			}
			continue
		}
		if i == 2 {
			if (nums[0] + nums[2]) > nums[1] {
				dp[2] = nums[0] + nums[2]
			} else {
				dp[2] = nums[1]
			}
			max = dp[i]
			continue
		}
		if (nums[i] + dp[i-2]) > (nums[i-1] + dp[i-3]) {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = nums[i-1] + dp[i-3]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
