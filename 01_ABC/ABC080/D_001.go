package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, 100001)
	var s, t, c int
	for i := 0; i < n; i++ {
		fmt.Scan(&s, &t, &c)
		a[s-1]++
		a[t]--
	}

	s = 0
	for i := 0; i < 100000; i++ {
		a[i+1] += a[i]
	}
	for i := 0; i < 100001; i++ {
		s = max(s, a[i])
	}
	fmt.Println(min(s, m))
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
