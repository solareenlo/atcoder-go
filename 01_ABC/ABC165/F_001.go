package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 1 << 60

var (
	N   int
	a   = make([]int, 0)
	ans = [200005]int{}
	dp  = make([]int, 0)
	cnt int
	G   = make([][]int, 200005)
)

func DFS(v int) {
	pos := lowerBound(dp, a[v])
	if dp[pos] == INF {
		cnt++
	}
	n := dp[pos]
	dp[pos] = a[v]
	ans[v] = cnt
	for _, i := range G[v] {
		if ans[i] == 0 {
			DFS(i)
		}
	}
	if n == INF {
		cnt--
	}
	dp[pos] = n
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &N)
	a = make([]int, N)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	dp = make([]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = INF
	}
	cnt = 0
	DFS(0)
	for i := 0; i < N; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
