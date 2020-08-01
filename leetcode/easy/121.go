package easy

/**
121. 买卖股票的最佳时机
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxProfit(prices []int) int {
	lastMinPrice := MAX
	lastMaxDiff := 0
	for _, price := range prices {
		lastMaxDiff = MaxInt(lastMaxDiff, price - lastMinPrice)
		lastMinPrice = MinInt(price, lastMinPrice)
		//fmt.Printf("min: %d, diff: %d\n", lastMinPrice, lastMaxDiff)
	}
	//fmt.Println()
	return lastMaxDiff
}