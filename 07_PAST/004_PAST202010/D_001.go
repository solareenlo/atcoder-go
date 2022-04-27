package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	X := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == '#' {
			X = append(X, i)
		}
	}

	ans := X[0] + (n - 1 - X[len(X)-1])
	for i := 0; i < len(X)-1; i++ {
		ans = max(ans, X[i+1]-X[i]-1)
	}
	fmt.Println(X[0], ans-X[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
