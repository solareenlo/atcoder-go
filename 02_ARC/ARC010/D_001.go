package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N  int
	c  int
	id = [5010]int{}
	ok = [5010]bool{}
	p  = [5010]int{}
	g  = [5010][5010]bool{}
)

func dfs1(pos int) {
	ok[pos] = true
	for i := 0; i < N; i++ {
		if !ok[i] && g[pos][i] {
			dfs1(i)
		}
	}
	p[c] = pos
	c++
}

func dfs2(pos int) {
	for i := 0; i < N; i++ {
		if id[i] == -1 && g[i][pos] {
			id[i] = id[pos]
			dfs2(i)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	xa := make([]int, N)
	ya := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &xa[i], &ya[i])
	}

	var M int
	fmt.Fscan(in, &M)
	xb := make([]int, M)
	yb := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &xb[i], &yb[i])
	}

	dp := make([]int, 5010)
	for i := 0; i < N; i++ {
		dp[i] = 1 << 60
		for j := 0; j < M; j++ {
			dp[i] = min(dp[i], (xa[i]-xb[j])*(xa[i]-xb[j])+(ya[i]-yb[j])*(ya[i]-yb[j]))
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if (xa[i]-xa[j])*(xa[i]-xa[j])+(ya[i]-ya[j])*(ya[i]-ya[j]) < dp[i] {
				g[i][j] = true
			}
		}
	}

	for i := 0; i < N; i++ {
		if !ok[i] {
			dfs1(i)
		}
	}

	for i := range id {
		id[i] = -1
	}
	cur := 0
	for i := N - 1; i >= 0; i-- {
		if id[p[i]] == -1 {
			id[p[i]] = cur
			dfs2(p[i])
			cur++
		}
	}
	for i := 0; i < cur; i++ {
		ok[i] = false
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if id[i] != id[j] && g[i][j] {
				ok[id[j]] = true
			}
		}
	}

	ret := 0
	for i := 0; i < cur; i++ {
		if !ok[i] {
			ret++
		}
	}
	fmt.Println(ret)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
