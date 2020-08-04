package two_sum

func TwoSum(nums []int, target int) []int {
	numMapIndex := map[int]int{}

	for index, n := range nums {
		numMapIndex[n] = index
		another_num := target - n
		another_index, exist := numMapIndex[another_num]
		if exist {
			return []int{index, another_index}
		}
	}

	for index, n := range nums {
		another_num := target - n
		another_index, exist := numMapIndex[another_num]
		if exist {
			return []int{index, another_index}
		}
	}

	return []int{}
}
