package simple_and_midium

func dominantIndex(nums []int) int {
	maxNum := -1
	maxNumIdx := -1
	secMaxNum := -1
	for i, n := range nums {
		if maxNum < n {
			secMaxNum = maxNum
			maxNum = n
			maxNumIdx = i
		} else if n > secMaxNum {
			secMaxNum = n
		}
	}

	if secMaxNum*2 > maxNum {
		return -1
	}
	return maxNumIdx
}
