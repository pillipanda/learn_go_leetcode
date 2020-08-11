package find_minimum_in_rotated_sorted_array

func find_min(subSlice []int) int {
	if len(subSlice) == 1 {
		return subSlice[0]
	}

	mid := len(subSlice)/2
	theArray := make([]int, mid)
	leftArray := subSlice[:mid]
	rightArray := subSlice[mid:]

	if leftArray[len(leftArray)-1] < rightArray[len(rightArray)-1] {
		theArray = leftArray
	} else {
		theArray = rightArray
	}
	return find_min(theArray)
}

func FindMinInRotatedSortedArray(sortedArray []int) int {
	return find_min(sortedArray)
}