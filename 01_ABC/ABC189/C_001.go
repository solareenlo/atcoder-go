package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	res := 0
	for l := 0; l < n; l++ {
		x := a[l]
		for r := l; r < n; r++ {
			x = min(x, a[r])
			res = max(res, x*(r-l+1))
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
