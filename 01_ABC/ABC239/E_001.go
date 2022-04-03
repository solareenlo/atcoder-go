package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	ar  []int
	g   [][]int
	ans [][]int
)

func dfs(v, p int) {
	ans[v] = []int{ar[v]}
	for _, to := range g[v] {
		if to == p {
			continue
		}
		dfs(to, v)
		ans[v] = append(ans[v], ans[to]...)
	}
	if len(ans[v]) >= 20 {
		sort.Sort(sort.Reverse(sort.IntSlice(ans[v])))
		ans[v] = ans[v][0:20]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	ar = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &ar[i])
	}

	g = make([][]int, n+1)
	ans = make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	dfs(1, -1)

	for i := 0; i < q; i++ {
		var v, k int
		fmt.Fscan(in, &v, &k)
		if len(ans[v]) != 20 {
			sort.Sort(sort.Reverse(sort.IntSlice(ans[v])))
		}
		fmt.Fprintln(out, ans[v][k-1])
	}
}
