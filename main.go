package main

import (
exo1 "TP/Exo1"
exo2 "TP/Exo2"
exo3 "TP/Exo3"
exo4 "TP/Exo4"
exo5 "TP/Exo5"
"fmt"
)

func main() {
fmt.Println("Exercice 1 :")
fmt.Println(exo1.Ft_coin([]int{1, 2, 5}, 11))
fmt.Println(exo1.Ft_coin([]int{2}, 3))
fmt.Println(exo1.Ft_coin([]int{1}, 0))

fmt.Println("Exercice 2 :")
fmt.Println(exo2.Ft_missing([]int{3, 1, 2}))                   // resultat : 0
fmt.Println(exo2.Ft_missing([]int{0, 1}))                      // resultat : 2
fmt.Println(exo2.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // resultat : 8

fmt.Println("Exercice 3 :")
fmt.Println(exo3.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
// il suffit de d'enlever qu'un seul intervalle, [1,3]
fmt.Println(exo3.Ft_non_overlap([][]int{{1, 2}, {2, 3}}))         // resultat : 0
fmt.Println(exo3.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}})) // resultat : 2

fmt.Println("Exercice 4 :")
fmt.Println(exo4.Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
fmt.Println(exo4.Ft_profit([]int{7, 6, 4, 3, 1}))    // resultat : 0

fmt.Println("Exercice 5 :")
fmt.Println(exo5.Ft_max_substring("abcabcbb")) // resultat : 3
fmt.Println(exo5.Ft_max_substring("bbbbb"))    // resultat : 1


}