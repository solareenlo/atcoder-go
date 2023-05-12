package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200010

var u, v, c, par, p, d, ans [N]int
var e [N][]int
var b []int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &u[i], &v[i], &c[i])
		u[i]--
		v[i]--
		if c[i] == 1 {
			e[u[i]] = append(e[u[i]], i)
			e[v[i]] = append(e[v[i]], i)
		}
		ans[i] = -1
	}
	dfs(0, -1)
	z := 0
	for i := 0; i < n; i++ {
		par[i] = i
	}
	for i := 0; i < m; i++ {
		if c[i] == 1 {
			if ans[i] == -1 {
				x := find(u[i])
				y := find(v[i])
				if d[x] < d[y] {
					par[y] = x
				} else {
					par[x] = y
				}
				ans[i] = z
				z++
			}
		} else {
			b = make([]int, 0)
			x := u[i]
			y := v[i]
			for {
				x = find(x)
				y = find(y)
				if x == y {
					break
				}
				if d[x] >= d[y] {
					unite(x)
				}
				if d[x] <= d[y] {
					unite(y)
				}
			}
			sort.Ints(b)
			for j := 0; j < len(b); j++ {
				ans[b[j]] = z
				z++
			}
			ans[i] = z
			z++
		}
	}
	for i := 0; i < m; i++ {
		if i == m-1 {
			fmt.Println(ans[i] + 1)
		} else {
			fmt.Printf("%d ", ans[i]+1)
		}
	}
}

func dfs(x, pv int) {
	for i := 0; i < len(e[x]); i++ {
		z := e[x][i]
		y := u[z] + v[z] - x
		if y == pv {
			continue
		}
		d[y] = d[x] + 1
		p[y] = z
		dfs(y, x)
	}
}

func find(x int) int {
	if par[x] == x {
		return x
	}
	par[x] = find(par[x])
	return par[x]
}

func unite(x int) {
	z := p[x]
	y := u[z] + v[z] - x
	b = append(b, z)
	par[find(x)] = find(y)
}
