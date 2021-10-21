package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct{ x, y, z int }

var (
	h   [100005]int
	de  [100005]int
	ans [100005]int
	c   [100005]int
	co  [100005]int
	f   [100005][20]int
	v   = make([][]node, 100005)
	q   = make([][]node, 100005)
)

func get(p int) {
	h[p] = h[f[p][0]] + 1
	for _, u := range v[p] {
		if u.x != f[p][0] {
			f[u.x][0] = p
			de[u.x] = de[p] + u.z
			get(u.x)
		}
	}
}

func lca(x, y int) int {
	if h[x] > h[y] {
		x, y = y, x
	}
	t := h[y] - h[x]
	for i := 0; i < 18; i++ {
		if t&(1<<i) != 0 {
			y = f[y][i]
		}
	}
	if x == y {
		return x
	}
	for i := 18; i >= 0; i-- {
		if f[x][i] != f[y][i] {
			x = f[x][i]
			y = f[y][i]
		}
	}
	return f[x][0]
}

func dfs(p int) {
	for _, u := range q[p] {
		if u.x > 0 {
			ans[u.x] += de[p] + c[u.y]*u.z - co[u.y]
		} else {
			ans[-u.x] -= 2 * (de[p] + c[u.y]*u.z - co[u.y])
		}
	}
	for _, u := range v[p] {
		if u.x != f[p][0] {
			c[u.y]++
			co[u.y] += u.z
			dfs(u.x)
			c[u.y]--
			co[u.y] -= u.z
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n-1; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		v[a] = append(v[a], node{b, c, d})
		v[b] = append(v[b], node{a, c, d})
	}
	get(1)
	for i := 1; i < 18; i++ {
		for j := 1; j < n+1; j++ {
			f[j][i] = f[f[j][i-1]][i-1]
		}
	}
	for i := 1; i < m+1; i++ {
		var x, y, u, v int
		fmt.Fscan(in, &u, &v, &x, &y)
		q[x] = append(q[x], node{i, u, v})
		q[y] = append(q[y], node{i, u, v})
		l := lca(x, y)
		q[l] = append(q[l], node{-i, u, v})
	}
	dfs(1)
	for i := 1; i < m+1; i++ {
		fmt.Println(ans[i])
	}
}
