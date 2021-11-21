package main

import "fmt"

var (
	n   int
	a       = [30]int{}
	res int = 1 << 60
)

func dfs(x, s1, s2 int) {
	if x == n {
		res = min(res, s1^s2)
		return
	}
	dfs(x+1, s1|a[x], s2)
	dfs(x+1, a[x], s1^s2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	dfs(0, 0, 0)

	fmt.Println(res)
}
