package simple_and_midium

// 1 <= prices.length <= 10^5

// maxProfit 暴力遍历
// prices 的最大长度10万，这么做会超时
func maxProfit(prices []int) int {
	rstProfit := 0 // record current max profit
	for i := len(prices) - 1; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			profit := prices[i] - prices[j]
			if profit > rstProfit {
				rstProfit = profit
			}
		}
	}

	return rstProfit
}

// maxProfit2 从后往前看，如果后一天的价格比前一天高，则把前一天的价格改为后一天的价格，记录这次的利润
func maxProfit2(prices []int) int {
	rstProfit := 0
	for i := len(prices) - 1; i >= 1; i-- {
		if prices[i] >= prices[i-1] {
			profit := prices[i] - prices[i-1]
			if rstProfit < profit {
				rstProfit = profit
			}
			// 价格前移
			prices[i-1] = prices[i]
		}
	}

	return rstProfit
}
