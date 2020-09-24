package decode_ways

func DecodeWays(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	sIntSlice := make([]int, len(s))
	for index, c := range s {
		intValue := int(c - '0')
		if intValue == 0 && (sIntSlice[index-1] > 2 || sIntSlice[index-1] == 0) {
			return 0
		}
		sIntSlice[index] = intValue
	}
	if sIntSlice[len(s)-1] == 0 && sIntSlice[len(s)-2] > 2 {
		return 0
	}

	dp := make([]int, len(sIntSlice))
	dp[0] = 1
	if sIntSlice[0] > 2 || (sIntSlice[0] == 2 && sIntSlice[1] > 6) || sIntSlice[1] == 0 {
		dp[1] = 1
	} else {
		dp[1] = 2
	}
	if len(sIntSlice) == 2 {
		return dp[1]
	}

	for i := 2; i < len(s); i++ {
		if sIntSlice[i-1] == 0 || sIntSlice[i] == 0 { // 若前面的为0则此数只能自己作为单数存在（且前数必须与前前数组【此情况dp[i-1]==dp[i-2]】）、若此数为0则只能和前数组(故等于dp[i-2])
			dp[i] = dp[i-2]
			continue
		}
		if sIntSlice[i-1] > 2 || (sIntSlice[i-1] == 2 && sIntSlice[i] > 6) { // 前一个数不能和自己进行组合
			dp[i] = dp[i-1]
		} else { // 演变为爬楼问题、对于i这个前面既可以是1步(dp[i-1])、也可以是2步(dp[i-2])
			dp[i] = dp[i-2] + dp[i-1]
		}
	}
	return dp[len(s)-1]
}
