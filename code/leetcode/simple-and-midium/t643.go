package simple_and_midium

func findMaxAverage(nums []int, k int) float64 {
	lastSum := 0
	for i := 0; i < k; i++ {
		lastSum += nums[i]
	}
	maxAverage := float64(lastSum) / float64(k)

	for i := 1; i+k-1 < len(nums); i++ {
		sum := lastSum - nums[i-1] + nums[i+k-1]
		avg := float64(sum) / float64(k)
		if maxAverage < avg {
			maxAverage = avg
		}
		lastSum = sum
	}

	return maxAverage
}
