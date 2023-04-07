package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	m := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		m = max(m, s[i]*a)
	}
	fmt.Println(m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
