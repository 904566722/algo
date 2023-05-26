package simple_and_midium

// 做一个假设：提前把需要消耗的汽油从加油站中抽掉
//          看作每段路不需要消耗汽油
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	for i := 0; i < n; i++ {
		gas[i] -= cost[i]
		sum += gas[i]
	}
	if sum < 0 {
		return -1
	}
	for i := 0; i < n; i++ {
		if gas[i] < 0 {
			continue
		}
		cnt := n - 1
		tmp := gas[i]
		nextId := (i+1)%n
		for cnt > 0 {
			tmp += gas[nextId]
			if tmp < 0 {
				break
			}
			cnt--
			nextId = (nextId+1)%n
		}
		if cnt == 0 {
			return i
		}
	}
	return -1
}