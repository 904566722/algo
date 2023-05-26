package simple_and_midium

func containsNearbyDuplicate(nums []int, k int) bool {
	rst := false
	idx := make(map[int]int) // k:v --> nums[i]:index
	for index, num := range nums {
		if firstIdx, ok := idx[num]; !ok {
			idx[num] = index
		} else {
			long := index - firstIdx
			if long <= k {
				rst = true
				break
			} else {
				// 更新下标为更大的，确保找到的两个相同的数是更近的
				idx[num] = index
			}
		}
	}

	return rst
}

// containsNearbyDuplicate2
// 滑动窗口，问题可以转化为：k宽度的窗口内是否存在重复元素
func containsNearbyDuplicate2(nums []int, k int) bool {
	set := make(map[int]bool)
	for j := 0; j < len(nums); j++ {
		if j > k {
			delete(set, nums[j-k-1])
		}

		if _, ok := set[nums[j]]; ok {
			return true
		} else {
			set[nums[j]] = true
		}
	}
	return false
}
