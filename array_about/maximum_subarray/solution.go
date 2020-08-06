package maximum_subarray

import (
	"fmt"
)

type MaxRecord struct {
	value int
	leftIndex int
	rightIndex int
}

func (max_rec *MaxRecord) compare (left_index, right_index, val int) {
	if val > max_rec.value {
		max_rec.value = val
		max_rec.leftIndex = left_index
		max_rec.rightIndex = right_index
	}
}

func MaxmiumSubArray(ipt []int) int {
	var sumIpt int
	for _, val := range ipt {
		sumIpt += val
	}

	maxRecord := &MaxRecord{value: sumIpt, leftIndex: 0, rightIndex: len(ipt) - 1}
	leftMove, rightMove := 0, 0
	for leftIndex, rightIndex := 0, len(ipt) - 1; leftIndex < rightIndex; {
		if leftMove <= rightMove {
			sumIpt -= ipt[leftIndex]
			leftIndex += 1
			leftMove += 1
		} else {
			sumIpt -= ipt[rightIndex]
			rightIndex -= 1
			rightMove += 1
		}
		maxRecord.compare(leftIndex, rightIndex, sumIpt)
	}

	fmt.Printf("between left index %d and right index %d create maxsub value %d",
		maxRecord.leftIndex, maxRecord.rightIndex, maxRecord.value)

	return maxRecord.value
}