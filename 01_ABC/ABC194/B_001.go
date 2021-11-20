package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}

	res := 1000001
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				res = min(res, a[i]+b[j])
			} else {
				res = min(res, max(a[i], b[j]))
			}
		}
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
