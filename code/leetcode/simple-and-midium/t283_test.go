package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	a := assert.New(t)
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums1)
	a.Equal([]int{1, 3, 12, 0, 0}, nums1)
	nums2 := []int{0}
	moveZeroes(nums2)
	a.Equal([]int{0}, nums2)
	nums3 := []int{1}
	moveZeroes(nums3)
	a.Equal([]int{1}, nums3)
	nums4 := []int{1, 1}
	moveZeroes(nums4)
	a.Equal([]int{1, 1}, nums4)
}

func TestMoveZeroes2(t *testing.T) {
	a := assert.New(t)
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes2(nums1)
	a.Equal([]int{1, 3, 12, 0, 0}, nums1)
	nums2 := []int{0}
	moveZeroes2(nums2)
	a.Equal([]int{0}, nums2)
	nums3 := []int{1}
	moveZeroes2(nums3)
	a.Equal([]int{1}, nums3)
	nums4 := []int{1, 1}
	moveZeroes2(nums4)
	a.Equal([]int{1, 1}, nums4)
}
