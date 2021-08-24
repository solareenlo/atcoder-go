package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	c := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
		dp[i] = int(1e18)
	}
	for i := 0; i < n; i++ {
		l, r := -1, n-1
		for r-l > 1 {
			m := (l + r) / 2
			if c[i] < dp[m] {
				r = m
			} else {
				l = m
			}
		}
		dp[r] = c[i]
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if dp[i] < int(1e18) {
			cnt++
		}
	}
	fmt.Println(n - cnt)
}
