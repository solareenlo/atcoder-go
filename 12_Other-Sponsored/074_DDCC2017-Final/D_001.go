package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

type tuple struct {
	x, y, z int
}

var n int
var g [55][]int
var dp map[tuple]int

func dfs(s, prev, pprev int) int {
	if s == (1<<n)-1 {
		return 1
	}
	for x := 0; x < n; x++ {
		for _, y := range g[x] {
			if x == prev || x == pprev {
				continue
			}
			if y == prev || y == pprev {
				continue
			}
			used1 := (s & (1 << x)) != 0
			used2 := (s & (1 << y)) != 0
			if used1 != used2 {
				return 0
			}
		}
	}
	key := tuple{s, prev, pprev}
	if dp[key] != 0 {
		return dp[key]
	}
	sum := 0
	for x := 0; x < n; x++ {
		if (s & (1 << x)) != 0 {
			continue
		}
		sum += dfs(s|(1<<x), x, prev)
	}
	dp[key] = sum % MOD
	return dp[key]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	dp = make(map[tuple]int)

	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	for i := 0; i < n; i++ {
		if len(g[i]) > 4 {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(dfs(0, -1, -1))
}
