package simple_and_midium

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	getCnt := func(nums []int) map[int]int {
		cnt := map[int]int{}
		for _, v := range nums {
			cnt[v] = 1
		}
		return cnt
	}
	cnt1, cnt2, cnt3 := getCnt(nums1), getCnt(nums2), getCnt(nums3)

	var ans []int
	for i := 0; i <= 100; i++ {
		if cnt1[i] + cnt2[i] + cnt3[i] >= 2 {
			ans = append(ans, i)
		}
	}
	return ans
}
