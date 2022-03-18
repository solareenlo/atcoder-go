package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200005

var (
	eb  = make([][]int, N)
	er  = make([][]int, N)
	fa  = make([]int, N)
	dep = make([]int, N)
	m   int
)

func dfs(x int) {
	for _, v := range eb[x] {
		if v != fa[x] {
			fa[v] = x
			dep[v] = dep[x] + 1
			dfs(v)
		}
	}
}

func check(u, v int) bool {
	if dep[u] < dep[v] {
		u, v = v, u
	}
	d := 1
	for d <= 2 && u != v {
		if dep[u] < dep[v] {
			v = fa[v]
		} else {
			u = fa[u]
		}
		d++
	}
	return u == v
}

func DFS(x, fa, k int) {
	m = max(m, dep[x])
	if dep[x] <= k {
		return
	}
	for _, v := range er[x] {
		if v != fa {
			if !check(x, v) {
				fmt.Println(-1)
				os.Exit(0)
			}
			DFS(v, x, k+1)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, U, R int
	fmt.Fscan(in, &n, &U, &R)

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		er[u] = append(er[u], v)
		er[v] = append(er[v], u)
	}

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		eb[u] = append(eb[u], v)
		eb[v] = append(eb[v], u)
	}

	dfs(R)
	DFS(U, 0, 0)

	fmt.Println(m << 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
