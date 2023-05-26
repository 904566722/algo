package simple_and_midium

import "sort"

func threeSum(nums []int) [][]int {
	// 先将数组从小到达排序
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int

	lastA := nums[0] - 1
	for i := 0; i < n; i++ {
		// 如果第一个数跟上一次循环的数一致，跳过该循环
		if nums[i] == lastA {
			continue
		}
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] + nums[j] + nums[k] == 0 {
					elem := []int{nums[i], nums[j], nums[k]}
					ans = append(ans, elem)
					break
				}
			}
		}

		lastA = nums[i]
	}
	return ans
}

func threeSum2(nums []int) [][]int {
	// 先将数组从小到达排序
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int

	lastA := nums[0] - 1
	for i := 0; i < n; i++ {
		// 如果第一个数跟上一次循环的数一致，跳过该循环
		cIdx := n - 1
		if nums[i] == lastA {
			continue
		}
		// 第二个数在往后寻找的时候
		// 如果下一个数跟上一个数相同，也应该跳过，否则会出现重复的三元组
		lastB := nums[0] - 1
		for j := i + 1; j < cIdx; j++ {
			if lastB == nums[j] {
				continue
			}
			for ; cIdx > j; cIdx-- {
				sum := nums[i] + nums[j] + nums[cIdx]

				if sum < 0 {
					// 如果和小于0，不必继往前遍历找第三个数
					// 让第二个数增大，才有可能让 sum 为零
					break
				} else if sum == 0 {
					elem := []int{nums[i], nums[j], nums[cIdx]}
					ans = append(ans, elem)
					break
				}
			}
			lastB = nums[j]
		}

		lastA = nums[i]
	}
	return ans
}