package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sub := make([][]int, n)
	var t int
	for i := 1; i < n; i++ {
		fmt.Scan(&t)
		sub[t-1] = append(sub[t-1], i)
	}

	var dfs func(int) int
	dfs = func(id int) int {
		if len(sub[id]) == 0 {
			return 1
		}
		maxi, mini := 0, int(1e9)
		for _, v := range sub[id] {
			p := dfs(v)
			maxi = max(maxi, p)
			mini = min(mini, p)
		}
		return maxi + mini + 1
	}
	fmt.Println(dfs(0))
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
