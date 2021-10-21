package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	for i := range a {
		fmt.Scan(&a[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		var b int
		fmt.Scan(&b)
		c := a[i] + a[i+1]
		res += min(b, c)
		a[i+1] = min(a[i+1], max(c-b, 0))
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
