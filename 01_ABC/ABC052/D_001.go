package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}

	res := 0
	for i := 0; i < n-1; i++ {
		res += min((x[i+1]-x[i])*a, b)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
