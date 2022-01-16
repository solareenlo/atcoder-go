package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	dp   = [5001][5001]int{}
	tree = make([][]int, 5000)
	sz   = [5001]int{}
	s    = [5001]int{}
)

func dfs(idx int) {
	dp[idx][0] = 0
	dp[idx][1] = s[idx]
	sz[idx] = 1
	for _, to := range tree[idx] {
		dfs(to)
		for i := sz[idx]; i > 0; i-- {
			for j := sz[to]; j >= 0; j-- {
				dp[idx][i+j] = min(dp[idx][i+j], dp[idx][i]+dp[to][j])
			}
		}
		sz[idx] += sz[to]
	}
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	all := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		all += s[i]
	}

	for i := 1; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		tree[a] = append(tree[a], b)
	}

	var m int
	fmt.Fscan(in, &m)

	t := make([]int, m+1)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &t[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(t)))

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
	}
	dfs(0)

	res := 0
	latte := 0
	lim := min(n, m)
	for i := 0; i < lim+1; i++ {
		res = max(res, all+latte-dp[0][i])
		latte += t[i]
	}
	fmt.Println(res)
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
