package final_prices_with_a_special_discount_in_a_shop

func FinalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	copy(ans, prices)
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				ans[i] = prices[i] - prices[j]
				break
			}
		}
	}
	return ans
}
