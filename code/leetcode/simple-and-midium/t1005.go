package simple_and_midium

import "math"

// 思路1
// m > 0
// 		k 奇数 rst = sum - min
// 		k 偶数 rst = sum
// m == 0	rst = sum
// m < = 把数组的最后一个 *-1
func largestSumAfterKNegations(nums []int, k int) int {
	sum := 0
	for i, n := 0, len(nums); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		sum += nums[i]
	}
	sum += nums[len(nums)-1]
	m := nums[0]

	if m == 0 || k == 0 {
		return sum
	} else if m > 0 {
		return sum - (m * (k % 2) * 2)
	}

	// m < 0
	ktmp := k
	for ; ktmp > 0 && m < 0 && k-ktmp < len(nums); ktmp-- {
		nums[k-ktmp] *= -1
		if k-ktmp+1 >= len(nums) {
			ktmp--
			break
		}
		m = nums[k-ktmp+1]
	}
	return largestSumAfterKNegations(nums, ktmp)
}

func largestSumAfterKNegations2(nums []int, k int) int {
	sum := 0
	m := math.MaxInt
	for i, n := 0, len(nums); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

		// nums[i] 为每次找到的最小值 m
		// 如果 m < 0 && k > 0 ，直接把这个最小的值变为整数
		if nums[i] < 0 && k > 0 {
			sum += nums[i] * -1
			m = min(m, nums[i]*-1)
			k--
		} else {
			sum += nums[i]
			m = min(m, nums[i])
		}
	}
	sum += nums[len(nums)-1]
	m = min(m, nums[len(nums)-1])

	return sum - (m * (k % 2) * 2)
}
