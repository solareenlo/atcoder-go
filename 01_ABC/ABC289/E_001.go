package main

import "fmt"

const N = 2005

type node struct {
	x1, x2, step int
}

var stk [N * N]node
var vis [N][N]bool
var color [N]int
var tail, head int
var g [N][]int

func Record(x1, x2, step int) {
	if vis[x1][x2] || color[x1] == color[x2] {
		return
	}
	vis[x1][x2] = true
	tail++
	stk[tail] = node{x1, x2, step}
}

func Solve() {
	var n, m int
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&color[i])
		g[i] = make([]int, 0)
		for j := 1; j <= n; j++ {
			vis[i][j] = false
		}
	}
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	head = 1
	tail = 0
	Record(1, n, 0)
	for head <= tail {
		q := stk[head]
		head++
		if q.x1 == n && q.x2 == 1 {
			fmt.Println(q.step)
			return
		}
		for i := 0; i < len(g[q.x1]); i++ {
			for j := 0; j < len(g[q.x2]); j++ {
				Record(g[q.x1][i], g[q.x2][j], q.step+1)
			}
		}
	}
	fmt.Println(-1)
	return
}

func main() {
	head = 1
	var T int
	fmt.Scan(&T)
	for T > 0 {
		T--
		Solve()
	}
}
