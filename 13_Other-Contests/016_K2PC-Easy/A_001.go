package main

import "fmt"

func main() {
	var a, b, c, N int
	fmt.Scan(&a, &b, &c, &N)
	fmt.Println(max(N-a, 0), max(2*N-b, 0), max(3*N-c, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
