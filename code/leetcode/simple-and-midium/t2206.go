package simple_and_midium

import (
	"fmt"
	"sort"
)

func divideArray(nums []int) bool {
	appeared := make([]bool, 501)
	for _, num := range nums {
		appeared[num] = !appeared[num]
	}
	for _, a := range appeared {
		if a {
			return false
		}
	}
	return true
}

func divideArray2(nums []int) bool {
	if len(nums) % 2 == 1 {
		return false
	}
	flg := nums[0]
	for _, n := range nums[1:] {
		flg ^= n
	}
	return flg == 0
}

func aa(nums []int)  {
	sort.Ints(nums)
	fmt.Println(nums)
}