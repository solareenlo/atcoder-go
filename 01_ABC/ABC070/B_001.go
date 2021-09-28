package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	mini := max(a, c)
	maxi := min(b, d)
	if mini >= maxi {
		fmt.Println(0)
	} else {
		fmt.Println(maxi - mini)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
