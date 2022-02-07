package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	s := make([]int, n+1)
	s[0] = a[0]
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + a[i]
	}

	res := int(1e16)
	for i := 0; i < n-1; i++ {
		res = min(res, abs(s[i]-(s[n-1]-s[i])))
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
