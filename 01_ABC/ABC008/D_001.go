package main

import "fmt"

type xy struct {
	xl, yl, xr, yr int
}

var h, w, n int
var a, b [30]int
var dp = map[xy]int{}

func dfs(xl, yl, xr, yr int) int {
	t := xy{xl: xl, yl: yl, xr: xr, yr: yr}
	if _, ok := dp[t]; ok {
		return dp[t]
	}
	res := 0
	for i := 0; i < n; i++ {
		if xl <= a[i] && a[i] < xr && yl <= b[i] && b[i] < yr {
			res = max(res, dfs(xl, yl, a[i], b[i])+dfs(a[i]+1, yl, xr, b[i])+dfs(xl, b[i]+1, a[i], yr)+dfs(a[i]+1, b[i]+1, xr, yr)+(xr-xl)+(yr-yl)-1)
		}
	}
	dp[t] = res
	return dp[t]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&h, &w, &n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
		a[i]--
		b[i]--
	}
	fmt.Println(dfs(0, 0, h, w))
}
