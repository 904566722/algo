package simple_and_midium

func minCostClimbingStairs(cost []int) int {
	pay := make([]int, len(cost)+1) // pay[i] 表示到达第i层楼梯需要第最小费用
	pay[1] = 0
	for i := 1; i < len(cost); i++ {
		pay[i+1] = min(pay[i]+cost[i], pay[i-1]+cost[i-1])
	}
	return pay[len(cost)]
}
