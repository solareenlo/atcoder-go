package main

import "fmt"

func main() {
	const INF = int(2e18)
	var n, m int
	fmt.Scan(&n, &m)
	ans := INF
	for i := 1; i <= n; i++ {
		x := (m + i - 1) / i
		if x <= n {
			ans = min(ans, i*x)
		}
		if i > x {
			break
		}
	}
	if ans == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
