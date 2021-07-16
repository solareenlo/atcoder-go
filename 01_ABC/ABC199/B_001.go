package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var x int
	l := 0
	r := int(1e9)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		l = max(l, x)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		r = min(r, x)
	}
	fmt.Println(max(0, r-l+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
