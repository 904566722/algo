package simple_and_midium

func pivotIndex(nums []int) int {
	sum := -nums[0]
	for _, n := range nums {
		sum += n
	}
	leftSum, rightSum := 0, sum
	i := 0
	for i < len(nums)-1 && leftSum != rightSum {
		i++
		leftSum += nums[i-1]
		rightSum -= nums[i]
	}

	if leftSum == rightSum {
		return i
	}
	return -1
}

//
func pivotIndex2(nums []int) int {
	// 求总和 total，设左侧和 sum，则右侧和 total-sum-nums[i]
	total := 0
	for _, n := range nums {
		total += n
	}

	sum := 0
	for i, n := range nums {
		if total-sum-n == sum {
			return i
		}
		sum += n
	}
	return -1
}
