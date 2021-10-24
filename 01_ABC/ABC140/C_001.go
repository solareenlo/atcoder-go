package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	b := make([]int, n-1)
	for i := range b {
		fmt.Scan(&b[i])
	}

	res := 0
	for i := 0; i < n-2; i++ {
		res += min(b[i], b[i+1])
	}

	fmt.Println(res + b[0] + b[n-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
