package exo2

func Ft_missing(num []int) int {
	n := len(num) // La longueur du tableau correspond à 'n'

	expectedSum := n * (n + 1) / 2

	actualSum := 0
	for _, value := range num {
		actualSum += value
	}

	return expectedSum - actualSum
}
