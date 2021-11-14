package main

import "fmt"

var (
	n, t, res int
	a         = [50]int{}
	sum       = [50]int{}
)

func dfs(w, s int) {
	if s == t {
		res = s
		return
	}
	if res == t {
		return
	}
	if s+sum[w] <= res {
		return
	}
	if w > n {
		res = s
		return
	}
	if s+a[w] <= t {
		dfs(w+1, s+a[w])
	} else {
		res = max(res, s)
	}
	dfs(w+1, s)
}

func main() {
	fmt.Scan(&n, &t)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&a[i])
	}
	for i := n; i >= 1; i-- {
		sum[i] = sum[i+1] + a[i]
	}
	dfs(1, 0)
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
