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

func FinalPrices2(prices []int) []int {
	ans := make([]int, len(prices))
	st := make([]int, len(prices))
	p := -1
	for i := 0; i < len(prices); i++ {
		for p >= 0 && (prices[st[p]] >= prices[i]) {
			ans[st[p]] = prices[st[p]] - prices[i]
			p--
		}
		p++
		st[p] = i
		ans[i] = prices[i]
	}
	return ans
}

func FinalPrices3(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{0}
	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		for len(st) > 1 && st[len(st)-1] > p {
			st = st[:len(st)-1]
		}
		ans[i] = p - st[len(st)-1]
		st = append(st, p)
	}
	return ans
}
