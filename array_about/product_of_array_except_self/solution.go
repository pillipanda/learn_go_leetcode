package product_of_array_except_self

func productOfArrayExceptSelf(ipt []int) []int{
	var leftAcc int = 1
	leftNum := make([]int, len(ipt), len(ipt))
	for index, val := range ipt {
		leftNum[index] = leftAcc
		leftAcc *= val
	}

	var rightAcc int = 1
	rightNum := make([]int, len(ipt), len(ipt))
	for index := len(ipt) - 1; index >= 0; index-- {
		rightNum[index] = rightAcc
		rightAcc *= ipt[index]
	}

	opt := make([]int, len(ipt), len(ipt))
	for index := 0; index < len(ipt); index ++ {
		opt[index] = leftNum[index] * rightNum[index]
	}

	return opt
}
