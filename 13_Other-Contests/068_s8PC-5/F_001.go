package main

import (
	"fmt"
	"sort"
)

const N = 55
const INF = 1e18

var n, m, q int
var a, b, x, y [N]int
var p []int
var f [N][N][N][N]int

func main() {
	fmt.Scan(&n, &m, &q)

	p = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		p[i] = i
	}
	tmp := p[1 : n+1]
	sort.Slice(tmp, func(x, y int) bool {
		return a[tmp[x]] < a[tmp[y]]
	})

	for i := 1; i <= n; i++ {
		b[p[i]] = i
	}
	for i := 1; i <= q; i++ {
		fmt.Scan(&x[i], &y[i])
	}
	for i := range f {
		for j := range f[i] {
			for k := range f[i][j] {
				for l := range f[i][j][k] {
					f[i][j][k][l] = -1
				}
			}
		}
	}
	fmt.Println(solve(m, 1, n, n+1))
}

func solve(t, l, r, lim int) int {
	if f[t][l][r][lim] != -1 {
		return f[t][l][r][lim]
	}
	if l > r {
		f[t][l][r][lim] = 0
		return f[t][l][r][lim]
	}
	f[t][l][r][lim] = INF
	p := 0
	for i := l; i <= r; i++ {
		if b[i] >= lim {
			continue
		}
		if p == 0 || b[i] > b[p] {
			p = i
		}
	}
	if t != 0 {
		f[t][l][r][lim] = min(f[t][l][r][lim], solve(t-1, l, r, b[p]))
	}
	if p == 0 {
		f[t][l][r][lim] = 0
		return f[t][l][r][lim]
	}
	cost := 0
	for i := 1; i <= q; i++ {
		if x[i] >= l && y[i] <= r && x[i] <= p && p <= y[i] {
			cost += a[p]
		}
	}
	for i := 0; i <= t; i++ {
		f[t][l][r][lim] = min(f[t][l][r][lim], solve(i, l, p-1, b[p])+solve(t-i, p+1, r, b[p])+cost)
	}
	return f[t][l][r][lim]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
