package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	deg = make([]int, 200200)
	x   = make([]int, 100100)
	g   = make([][]int, 200200)
	n   int
	m   int
)

func check(d int) bool {
	q := make([]int, 0)
	cdeg := make([]int, 200200)
	copy(cdeg, deg)
	dp := make([]int, 200200)
	for i := range dp {
		dp[i] = -1
	}
	for i := 0; i < n; i++ {
		if x[i] >= d {
			dp[i] = 1
			q = append(q, i)
			if cdeg[i] == 0 {
				dp[i+n] = 0
				q = append(q, i+n)
			}
		} else {
			dp[i+n] = 1
			q = append(q, i+n)
			if cdeg[i] == 0 {
				dp[i] = 0
				q = append(q, i)
			}
		}
	}
	for len(q) > 0 {
		now := q[0]
		q = q[1:]
		for _, nxt := range g[now] {
			if dp[nxt] != -1 {
				continue
			}
			cdeg[nxt]--
			if dp[now] == 0 {
				dp[nxt] = 1
				q = append(q, nxt)
			} else if dp[now] == 1 {
				if cdeg[nxt] == 0 {
					dp[nxt] = 0
					q = append(q, nxt)
				}
			}
		}
	}
	return dp[0] == 1
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i])
	}
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[v] = append(g[v], u+n)
		g[v+n] = append(g[v+n], u)
		deg[u]++
		deg[u+n]++
	}

	l, r := 0, 1<<60
	for r-l > 1 {
		mid := (r + l) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Println(l)
}
