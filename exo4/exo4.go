package exo4

func Ft_profit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > res {
				res = prices[j] - prices[i]
			}
		}
	}
	return res
}
