package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

var (
	p = make([]int, N)
	g = make([]map[int]bool, N)
)

func find(x int) int {
	if x == p[x] {
		return x
	}
	p[x] = find(p[x])
	return p[x]
}

type pair struct{ x, y int }

func merge(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	if len(g[u]) < len(g[v]) {
		u, v = v, u
	}
	que := map[pair]bool{}
	for x := range g[v] {
		if x == u {
			continue
		}
		delete(g[x], v)
		if _, ok := g[x][u]; ok {
			que[pair{x, u}] = true
		}
		g[u][x] = true
		g[x][u] = true
	}
	delete(g[u], v)
	g[v] = make(map[int]bool)
	p[v] = u
	for k := range que {
		merge(k.x, k.y)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < N; i++ {
		p[i] = i
	}

	for i := range g {
		g[i] = make(map[int]bool)
	}
	for i := 1; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a][b] = true
		g[b][a] = true

	}

	for i := 1; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a = find(a)
		b = find(b)
		if _, ok := g[a][b]; ok {
			merge(a, b)
		} else {
			g[a][b] = true
			g[b][a] = true
		}
	}

	c := find(1)
	for i := 2; i <= n; i++ {
		if c != find(i) {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
