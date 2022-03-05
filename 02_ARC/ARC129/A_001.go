package main

import "fmt"

func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)

	ans := 0
	for i := 0; i < 60; i++ {
		if (n>>i)&1 != 0 {
			ans += max(min(r+1, (1<<(i+1)))-max(l, 1<<i), 0)
		}
	}
	fmt.Println(ans)
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
