package simple_and_midium

import "leetcode/utils"

func findMaxConsecutiveOnes(nums []int) int {
	rst := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		cnt += nums[i]

		if nums[i] == 0 {
			rst = utils.Max(rst, cnt)
			cnt = 0
		}
	}
	rst = utils.Max(rst, cnt)
	return rst
}
