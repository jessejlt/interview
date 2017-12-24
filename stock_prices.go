package interview

// BuySellStocks determines the maximum gain possible given a list of stock prices for a given day.
func BuySellStocks(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	var (
		buy  = 0
		sell = 1

		profit, maxProfit int
	)

	for ; sell < len(prices); sell++ {

		if prices[buy] < prices[sell] {

			profit = intMax(profit, prices[sell]-prices[buy])
		} else {

			if profit > maxProfit {

				maxProfit = profit
			}

			profit = 0
			buy = sell
		}
	}

	return intMax(profit, maxProfit)
}
