package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n %= 30

	card := [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < n; i++ {
		j := i % 5
		card[j], card[j+1] = card[j+1], card[j]
	}

	for i := range card {
		fmt.Print(card[i])
	}
	fmt.Println()
}
