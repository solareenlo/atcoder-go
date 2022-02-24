package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	siz = [100100]int{}
	dp  = [100100]int{}
	g   = make([][]int, 100100)
)

func dfs(x int) {
	s := make([]int, 0)
	siz[x] = 1
	dp[x] = 1
	sum := 0
	for _, v := range g[x] {
		dfs(v)
		siz[x] += siz[v]
		if siz[v]&1 != 0 {
			s = append(s, dp[v])
		} else {
			if dp[v] < 0 {
				dp[x] += dp[v]
			} else {
				sum += dp[v]
			}
		}
	}
	sort.Ints(s)
	s = append(s, sum)
	for i := 0; i < len(s); i++ {
		if i&1 != 0 {
			dp[x] -= s[i]
		} else {
			dp[x] += s[i]
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 2; i <= n; i++ {
		var p int
		fmt.Fscan(in, &p)
		g[p] = append(g[p], i)
	}
	dfs(1)
	fmt.Println((n + dp[1]) / 2)
}
