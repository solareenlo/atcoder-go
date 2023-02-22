package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INF = 1e18
const INF2 = 1e19
const MAXN = 114514

var x, y, dis []float64
var hd []int
var tt int
var e []E
var cost float64
var vis []bool

type E struct {
	u, v, nx int
	w, c     float64
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	x = make([]float64, MAXN)
	y = make([]float64, MAXN)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	st := n*2 + 1
	en := n*2 + 2
	e = make([]E, MAXN)
	hd = make([]int, MAXN)
	dis = make([]float64, MAXN)
	vis = make([]bool, MAXN)
	tt = 1
	for i := 1; i <= n; i++ {
		ade(st, i, 1.0, 0)
		for j := 1; j <= n; j++ {
			ade(i, j+n, 1.0, dist(i, j))
		}
		ade(i+n, en, 1.0, 0)
	}
	mcmf(st, en)
	fmt.Println(cost)
}

func mcmf(s, t int) float64 {
	ans := 0.0
	for spfa(s, t) {
		var x float64
		for {
			x = dfs(s, t, INF)
			if x == 0 {
				break
			}
			ans += x
		}
	}
	return ans
}

func dfs(u, t int, flow float64) float64 {
	if u == t {
		return flow
	}
	vis[u] = true
	ans := 0.0
	for i := hd[u]; i > 0 && ans < flow; i = e[i].nx {
		v := e[i].v
		if !vis[v] && e[i].w != 0 && dis[v] == dis[u]+e[i].c {
			x := dfs(v, t, math.Min(e[i].w, flow-ans))
			if x != 0 {
				cost += x * e[i].c
				e[i].w -= x
				e[i^1].w += x
				ans += x
			}
		}
	}
	vis[u] = false
	return ans
}

func spfa(s, t int) bool {
	for i := range dis {
		dis[i] = INF2
	}
	q := make([]int, 0)
	q = append(q, s)
	dis[s] = 0
	vis[s] = true
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		vis[u] = false
		for i := hd[u]; i > 0; i = e[i].nx {
			v := e[i].v
			if e[i].w != 0 && dis[v] > dis[u]+e[i].c {
				dis[v] = dis[u] + e[i].c
				if !vis[v] {
					q = append(q, v)
					vis[v] = true
				}
			}
		}
	}
	return dis[t] != dis[0]
}

func dist(i, j int) float64 {
	return math.Sqrt((x[i]+x[j])*(x[i]+x[j])+(y[i]-y[j])*(y[i]-y[j])) / 2
}

func ade(u, v int, w, c float64) {
	tt++
	e[tt].u = u
	e[tt].v = v
	e[tt].w = w
	e[tt].c = c
	e[tt].nx = hd[u]
	hd[u] = tt
	tt++
	e[tt].u = v
	e[tt].v = u
	e[tt].w = 0
	e[tt].c = -c
	e[tt].nx = hd[v]
	hd[v] = tt
}
