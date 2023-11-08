package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, ans int
var a, b, d [50]int
var g []int

func get(s, h int) int {
	for i := 0; i < n; i++ {
		if ((s >> i) & 1) != 0 {
			h += b[i] - a[i]
		}
	}
	return h
}

func dfs(x, s, h, v int) {
	ans = max(ans, v)
	for i := x; i < len(g); i++ {
		ns := s | d[g[i]]
		nh := h + b[g[i]] - a[g[i]]
		nv := get(ns, nh)
		if nv > v {
			dfs(i+1, ns, nh, nv)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var c [50]int
	var u, v [100100]int

	var m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		if a[i] < b[i] {
			c[i] = 1
		} else {
			c[i] = 0
		}
	}
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &u[i], &v[i])
		u[i]--
		v[i]--
		if u[i] == v[i] && c[u[i]] != 0 {
			a[u[i]], b[u[i]] = b[u[i]], a[u[i]]
			c[u[i]] = 0
		}
	}
	s := 0
	for i := 1; i <= m; i++ {
		if c[u[i]] != 0 && c[v[i]] != 0 {
			s |= (1 << u[i])
			s |= (1 << v[i])
		} else if c[u[i]] != 0 {
			d[v[i]] |= (1 << u[i])
		} else if c[v[i]] != 0 {
			d[u[i]] |= (1 << v[i])
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		if c[i] == 0 {
			g = append(g, i)
		}
		sum += a[i] * 2
	}
	dfs(0, s, 0, get(s, 0))
	fmt.Println(float64(ans+sum) / 2.0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
