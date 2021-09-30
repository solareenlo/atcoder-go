package main

import "fmt"

func main() {
	var n, k, x int
	fmt.Scan(&n, &k)

	res := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		res += min(x, abs(x-k)) * 2
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
