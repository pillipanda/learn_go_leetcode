package container_with_most_water

func ContainerWithMostWater (ipt []int) int {
	var (
		leftIndex, maxWater int
		rightIndex int = len(ipt) - 1
	)

	for leftIndex < rightIndex {
		var thisWater int
		if ipt[rightIndex] > ipt[leftIndex] {
			thisWater = (rightIndex - leftIndex) * ipt[leftIndex]
			leftIndex += 1
			for leftIndex < rightIndex {
				if ipt[leftIndex] <= ipt[leftIndex - 1] {
					leftIndex += 1
				} else {
					break
				}
			}
		} else {
			thisWater = (rightIndex - leftIndex) * ipt[rightIndex]
			rightIndex -= 1
			for leftIndex < rightIndex {
				if ipt[rightIndex] <= ipt[rightIndex + 1] {
					rightIndex -= 1
				} else {
					break
				}
			}
		}
		if thisWater > maxWater {
			maxWater = thisWater
		}
	}
	return maxWater
}