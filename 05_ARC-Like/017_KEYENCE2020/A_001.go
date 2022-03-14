package main

import "fmt"

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)

	fmt.Println((n-1)/max(h, w) + 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
