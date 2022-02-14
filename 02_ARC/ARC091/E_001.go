package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if a+b > n+1 || a*b < n {
		fmt.Println(-1)
		return
	}

	for n > 0 {
		for i := n - min(a, n-b+1) + 1; i <= n; i++ {
			fmt.Print(i, " ")
		}
		n -= min(a, n-b+1)
		b--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
