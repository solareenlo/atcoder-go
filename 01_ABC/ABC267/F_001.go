package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 200001

var edge, query [][]int
var dis, node, ans [MAXN]int
var len int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	edge = make([][]int, MAXN)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		edge[u] = append(edge[u], v)
		edge[v] = append(edge[v], u)
	}
	var q int
	fmt.Fscan(in, &q)
	query = make([][]int, MAXN)
	for i := 1; i <= q; i++ {
		var u int
		fmt.Fscan(in, &u, &dis[i])
		query[u] = append(query[u], i)
	}
	for _, turn := range []int{0, 1, 2} {
		pos := node[len]
		len = 0
		if turn == 0 {
			pos = 1
		}
		dfs(pos, 0, 0)
	}

	for i := 1; i <= q; i++ {
		if ans[i] != 0 {
			fmt.Println(ans[i])
		} else {
			fmt.Println(-1)
		}
	}
}

func dfs(p, f, dep int) {
	node[dep] = p
	for _, q := range query[p] {
		if dep >= dis[q] {
			ans[q] = node[dep-dis[q]]
		}
	}
	len = max(len, dep)
	for _, v := range edge[p] {
		if v != f {
			dfs(v, p, dep+1)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
