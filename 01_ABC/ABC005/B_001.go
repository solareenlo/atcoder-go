package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n)
	res := 100
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		res = min(res, t)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
