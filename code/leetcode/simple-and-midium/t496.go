package simple_and_midium

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	childSet := make(map[int]int)
	rst := make([]int, len(nums1))
	for i, v := range nums1 {
		childSet[v] = i
	}

	for i, v := range nums2 {
		if idx, ok := childSet[v]; ok {
			if i >= len(nums2)-1 {
				rst[idx] = -1
			}
			for j := i + 1; j < len(nums2); j++ {
				if nums2[j] > v {
					rst[idx] = nums2[j]
					break
				}
				if j >= len(nums2)-1 {
					rst[idx] = -1
				}
			}
		}
	}

	return rst
}

// 使用栈来实现
type myStack struct {
	Nums []int
}

func (s *myStack) push(n int) {
	s.Nums = append(s.Nums, n)
}

func (s *myStack) pop() int {
	if s.isEmpty() {
		return 0
	}
	top := s.Nums[len(s.Nums)-1]
	s.Nums = s.Nums[:len(s.Nums)-1] // num[n:m] 的下标范围 [n,m)
	return top
}

func (s *myStack) top() int {
	return s.Nums[len(s.Nums)-1]
}

func (s *myStack) isEmpty() bool {
	if len(s.Nums) == 0 {
		return true
	}
	return false
}

// nextGreaterElement2 使用单调栈来实现
// 问题分解:1. 先处理 nums2，找到每个数对应的结果
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	var rst []int
	set := make(map[int]int)
	s := myStack{
		Nums: []int{},
	}
	for i := len(nums2) - 1; i >= 0; i-- {
		n := nums2[i]
		// 栈不空 且 n 比栈顶元素大，出栈
		for !s.isEmpty() && n > s.top() {
			s.pop()
		}
		// 栈空，或者 n 比栈顶元素小
		if s.isEmpty() {
			set[n] = -1
		} else {
			set[n] = s.top()
		}
		s.push(n)
	}

	for i := 0; i < len(nums1); i++ {
		rst = append(rst, set[nums1[i]])
	}
	return rst
}
