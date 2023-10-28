package main

import (
	"fmt"
	"sort"
)

const SIZE = 100005

var vec [SIZE][]int
var d1, d2 [SIZE]int
var dist [SIZE][2]int
var N int

func dfs(v, p int) {
	d1[v] = 0
	d2[v] = -1
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			dfs(to, v)
			vl := d1[to] + 1
			if vl > d1[v] {
				vl, d1[v] = d1[v], vl
			}
			if vl > d2[v] {
				vl, d2[v] = d2[v], vl
			}
		}
	}
}

func Make(v, p int) {
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			vl := d1[v] + 1
			if d1[v] == d1[to]+1 {
				vl = d2[v] + 1
			}
			if vl > d1[to] {
				vl, d1[to] = d1[to], vl
			}
			if vl > d2[to] {
				vl, d2[to] = d2[to], vl
			}
			Make(to, v)
		}
	}
}

func solve() {
	fmt.Scan(&N)
	n := N
	for i := 0; i < n; i++ {
		vec[i] = make([]int, 0)
	}
	for i := 0; i < n-1; i++ {
	}
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		vec[a] = append(vec[a], b)
		vec[b] = append(vec[b], a)
	}
	dfs(0, -1)
	Make(0, -1)
}

func main() {
	solve()
	n := N
	for i := 0; i < n; i++ {
		dist[i][0] = d1[i]
	}
	solve()
	m := N
	for i := 0; i < m; i++ {
		dist[i][1] = d1[i]
	}
	for i := 0; i < n; i++ {
		d1[i] = dist[i][0]
	}
	for i := 0; i < m; i++ {
		d2[i] = dist[i][1]
	}
	L := 0
	for i := 0; i < n; i++ {
		L = max(L, d1[i])
	}
	for i := 0; i < m; i++ {
		L = max(L, d2[i])
	}
	ret := 0
	vx := make([]int, 0)
	vy := make([]int, 0)
	for i := 0; i < n; i++ {
		vx = append(vx, d1[i])
	}
	for i := 0; i < m; i++ {
		vy = append(vy, d2[i])
	}
	sort.Ints(vx)
	sort.Ints(vy)
	to := len(vy)
	sum := 0
	for i := 0; i < len(vx); i++ {
		for to > 0 && vy[to-1]+vx[i]+1 >= L {
			to--
			sum += vy[to]
		}
		ret += sum + (vx[i]+1)*(m-to)
		ret += L * to
	}
	fmt.Println(ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
