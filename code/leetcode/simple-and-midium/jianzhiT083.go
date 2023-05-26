package simple_and_midium

import (
	"sort"
)

func permute(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	permuteNum := 1
	for i := n; i>=1 ; i-- {
		permuteNum *= i
	}

	var ans [][]int
	for permuteNum > 0 {
		// deep copy
		tmpNums := make([]int, len(nums))
		copy(tmpNums, nums)
		ans = append(ans, tmpNums)

		nextDicArr(nums)
		permuteNum--
	}
	return ans
}

func nextDicArr(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pos := -1
	for i := len(nums)-2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pos = i
			break
		}
	}
	// 没找到上坡（左值小于右值），说明已经是降序排列
	if pos == -1 {
		for i,n := 0,len(nums); i < n/2; i++ {
			nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
		}
		return
	}

	// 找到 pos，接着找右边大于该值的最小值（也就是第一个大于标记的值）
	biggerThanPos := -1
	for i := len(nums)-1; i >= pos+1; i-- {
		if nums[i] > nums[pos] {
			biggerThanPos = i
			break
		}
	}
	nums[pos], nums[biggerThanPos] = nums[biggerThanPos], nums[pos]

	// 剩下的倒序
	for i,n,cnt := pos+1, len(nums),0; i < (n + pos+1)/2; i++ {
		nums[i], nums[n-cnt-1] = nums[n-cnt-1], nums[i]
		cnt++
	}
	return
}