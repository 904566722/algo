package simple_and_midium

import "math"

func findMatrix(nums []int) [][]int {
	freq := make([]int, 201)
	maxFreq := math.MinInt
	for _, num := range nums {
		freq[num]++
		if maxFreq < freq[num] {
			maxFreq = freq[num]
		}
	}
	ans := make([][]int, maxFreq)
	for i := 1; i < len(freq); i++ {
		if freq[i] == 0 {
			continue
		}

		for row := 0; row < freq[i]; row++ {
			ans[row] = append(ans[row], i)
		}
	}
	return ans
}
