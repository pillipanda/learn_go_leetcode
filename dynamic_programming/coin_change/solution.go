package coin_change

import (
	"math"
)

func CoinChange(coins []int, target int) int {
	if target <= 0 || len(coins) <= 0 {
		return 0
	}
	dp := make([]int, target + 1)
	dp[0] = 0
	for i := 1; i <= target; i ++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if (i - coin) < 0 {
				continue
			}
			if dp[i] > (dp[i - coin] + 1) {
				dp[i] = dp[i - coin] + 1
			}
		}
	}
	if dp[target] == math.MaxInt32 {
		return -1
	}
	return dp[target]
}
