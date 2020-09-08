package word_break

func WordBreak(s string, wordDict []string) bool {
	dp := make([]int, len(s))
	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}

	var subWord string
	for index := 0; index < len(s); index++ {
		subWord = s[:index+1]
		_, exist := wordMap[subWord]
		if exist {
			dp[index] = 1
			continue
		}
		for already_index := 0; already_index < index; already_index++ {
			if dp[already_index] > 0 {
				tmp := subWord[already_index+1:]
				_, exist := wordMap[tmp]
				if exist {
					dp[index] = 1
				}
			}
		}
	}

	if dp[len(s)-1] > 0 {
		return true
	}

	return false
}
