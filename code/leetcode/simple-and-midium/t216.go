package simple_and_midium


// 剪枝条件：当前的和 > n | 当前的数量 > k
func CombinationSum3(k int, n int) [][]int {
	state := &State{
		sum: 0,
	}
	choices := make([]int, 9)
	for i := 1; i <= 9; i++ {
		choices[i-1] = i
	}
	var ans [][]int
	sumkHelper(state, n, k, choices, &ans)
	return ans
}

type State struct {
	sum int
	nums []int
}

func sumkHelper(state *State, n, k int, choices []int, ans *[][]int) {
	/*判断当前是否符合解的条件*/
	if state.sum == n && len(state.nums) == k {
		tmp := make([]int, k)
		copy(tmp, state.nums)
		*ans = append(*ans, tmp)
		return
	}

	for i, choice := range choices {
		/*剪枝*/
		if state.sum > n || len(state.nums) > k {
			return
		}
		// 由于 choices 从小到大排序的
		// 提前判断，如果做出选择已经大于 n，就没有继续循环的必要
		if state.sum + choice > n {
			break
		}

		// 尝试
		state.sum += choice
		state.nums = append(state.nums, choice)
		sumkHelper(state, n, k, choices[i+1:], ans)
		// 回退
		state.sum -= choice
		state.nums = state.nums[:len(state.nums)-1]
	}
}
