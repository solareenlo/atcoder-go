package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	fmt.Println(min(n*a, b))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
