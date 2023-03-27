package main

import (
	"fmt"
	"sort"
)

const INF = int(1e18)

var n, m int
var s [1005]int
var dp [1005][1005]int

func cost(a, b int) int {
	if b < a {
		return 0
	}
	if (b-a+1)%2 == 1 {
		if a == 1 {
			return s[b] - s[(a+b)/2] - s[(a+b)/2-1]
		}
		return s[b] - s[(a+b)/2] - s[(a+b)/2-1] + s[a-1]
	} else {
		return s[b] - 2*s[(a+b)/2] + s[a-1]
	}
}

func dfs(q int, l, r, a, b int) {
	x := (l + r) / 2
	p := -1
	for i := a; i <= min(b, x); i++ {
		if dp[q][x] > dp[q-1][i]+cost(i+1, x) {
			dp[q][x] = dp[q-1][i] + cost(i+1, x)
			p = i
		}
	}
	if x-1 >= l {
		dfs(q, l, x-1, a, p)
	}
	if x+1 <= r {
		dfs(q, x+1, r, p, b)
	}
}

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &s[i])
	}
	tmp := s[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	for i := 1; i <= n; i++ {
		s[i] += s[i-1]
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dfs(i, 1, n, 0, n)
	}
	fmt.Printf("%d\n", dp[m][n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
