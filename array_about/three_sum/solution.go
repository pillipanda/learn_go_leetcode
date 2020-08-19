package three_sum

import (
	"sort"
)

func ThreeSumBS(ipt []int) [][]int {
	result := make([][]int, 0, len(ipt))
	if len(ipt) < 3 {
		return result
	}

	NumMapCount := make(map[int]int, len(ipt))
	for _, num := range ipt {
		NumMapCount[num] += 1
	}

	uniSlice := make([]int, 0, len(NumMapCount))
	for k := range NumMapCount {
		uniSlice = append(uniSlice, k)
	}
	sort.Ints(uniSlice)

	for i := 0; i < len(uniSlice); i++ {
		value := uniSlice[i]
		if value == 0 && NumMapCount[value] >= 3 {
			result = append(result, []int{0, 0, 0})
			continue
		}

		for j := i + 1; j < len(uniSlice); j++ {
			if value * 2 + uniSlice[j] == 0 && NumMapCount[value] >= 2 {
				result = append(result, []int{value, value, uniSlice[j]})
			}

			if value + uniSlice[j] * 2 == 0 && NumMapCount[uniSlice[j]] >= 2 {
				result = append(result, []int{value, uniSlice[j], uniSlice[j]})
			}

			leftNum := -(value + uniSlice[j])
			if NumMapCount[leftNum] >= 1 && leftNum > uniSlice[j] {
				result = append(result, []int{value, uniSlice[j], leftNum})
			}
		}
 	}

	return result
}