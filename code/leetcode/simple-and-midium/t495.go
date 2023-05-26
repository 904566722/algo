package simple_and_midium

func findPoisonedDuration(timeSeries []int, duration int) int {
	sum := 0
	for i := 1; i < len(timeSeries); i++ {
		attachDuration := timeSeries[i] - timeSeries[i-1]
		if attachDuration <= duration-1 {
			sum += attachDuration
		} else {
			sum += duration
		}
	}
	// last attack
	sum += duration
	return sum
}
