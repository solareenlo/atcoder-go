package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

type pair struct{ x, y int }

var (
	G    = make([][]int, N)
	from = make([][]int, N)
	vis  = make([]int, N)
	mat  = make([]int, N)
	fa   = make([]int, N)
	son  = make([]int, N)
	ans  = make([]pair, N)
)

func DFS(x, tim int) bool {
	for _, v := range G[x] {
		if vis[v] != tim {
			vis[v] = tim
			if mat[v] == 0 || DFS(mat[v], tim) {
				mat[v] = x
				return true
			}
		}
	}
	return false
}

func find(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func build(u int) {
	for _, v := range from[u] {
		if ans[v].x == 0 {
			ans[v] = pair{u, son[v]}
			build(son[v])
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fa[i] = i
	}

	var t0 int
	for i := 1; i < n; i++ {
		var s int
		fmt.Fscan(in, &s)
		for j := 0; j < s; j++ {
			var t int
			fmt.Fscan(in, &t)
			G[i] = append(G[i], t)
			from[t] = append(from[t], i)
			if j == 0 {
				t0 = t
			} else {
				fa[find(t)] = find(t0)
			}
		}
	}

	for i := 1; i <= n; i++ {
		if find(i) != find(n) {
			fmt.Println(-1)
			return
		}
	}

	for i := 1; i < n; i++ {
		if !DFS(i, i) {
			fmt.Println(-1)
			return
		}
	}

	for i := 1; i <= n; i++ {
		son[mat[i]] = i
	}

	for i := 1; i <= n; i++ {
		if mat[i] == 0 {
			build(i)
			for j := 1; j < n; j++ {
				if ans[j].x == 0 {
					fmt.Println(-1)
					return
				}
			}
			for j := 1; j < n; j++ {
				fmt.Println(ans[j].x, ans[j].y)
			}
			return
		}
	}
}
