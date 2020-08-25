package climbing_stairs

func ClimbingStairs(n int) int {
	dp := [3]int{0, 1, 2}
	if n < len(dp) {
		return dp[n]
	}
	for i := 3; i <= n; i++ {
		dp[1], dp[2] = dp[2], dp[1] + dp[2]
	}
	return dp[2]
}
