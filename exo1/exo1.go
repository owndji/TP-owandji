package exo1

func Ft_coin(coins []int, amount int) int {
	if amount == 0 {
	return 0
	}
	
	count := 0
	value := 0
	for amount != 0 {
	for i := 0; i < len(coins); i++ {
	if value <= coins[i] && coins[i] <= amount {
	value = coins[i]
	}
	}
	if value == 0 {
	return -1
	} else if value <= amount {
	count++
	amount -= value
	value = 0
	}
	}
	
	return count
	}