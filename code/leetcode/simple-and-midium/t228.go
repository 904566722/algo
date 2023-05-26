package simple_and_midium

import (
	"fmt"
)

// summaryRanges
func summaryRanges(nums []int) []string {
	var rst []string
	if len(nums) == 0 {
		return []string{}
	}
	var start, pre, cur int = nums[0], nums[0], nums[0]
	for _, num := range nums[1:] {
		cur = num

		// 断层，计算前面的一个区间
		if cur != pre+1 {
			var rangeStr string
			if start != pre {
				rangeStr = fmt.Sprintf("%d->%d", start, pre)
			} else {
				rangeStr = fmt.Sprintf("%d", start)
			}
			rst = append(rst, rangeStr)

			start = cur
		}

		pre = cur
	}

	// 计算最后的一个区间
	if start == cur {
		rst = append(rst, fmt.Sprintf("%d", start))
	} else {
		rst = append(rst, fmt.Sprintf("%d->%d", start, cur))
	}

	return rst
}
