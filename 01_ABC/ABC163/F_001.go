package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	c   = [200005]int{}
	sum = [200005]int{}
	res = [200005]int{}
	sub = [200005]int{}
	g   = make([][]int, 200005)
)

func f(x int) int {
	return x * (x + 1) / 2
}

func dfs(u, fa int) {
	t := sum[c[u]]
	sub[u] = 1
	for _, v := range g[u] {
		sum[c[u]] = 0
		if v == fa {
			continue
		}
		dfs(v, u)
		sub[u] += sub[v]
		res[c[u]] -= f(sub[v] - sum[c[u]])
	}
	sum[c[u]] = t + sub[u]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &c[i])
	}

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	dfs(1, 1)

	for i := 1; i < n+1; i++ {
		fmt.Fprintln(out, res[i]+f(n)-f(n-sum[i]))
	}
}
