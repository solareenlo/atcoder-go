package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const nMax = 305
const INF = int(1e9)

var n, m int
var res int = INF
var g [nMax][]pair
var dis [nMax][nMax]int
var deg [nMax]int

func Ins(cur int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dis[i][j] = min(dis[i][j], dis[i][cur]+dis[cur][j])
		}
	}
}

func solve(l, r int) {
	if l == r {
		if deg[l] < 3 {
			return
		}
		for _, A := range g[l] {
			for _, B := range g[l] {
				if !A.lessThan(B) {
					continue
				}
				for _, C := range g[l] {
					if (!C.eq(A)) && (!C.eq(B)) {
						res = min(res, dis[A.y][B.y]+A.x+B.x+C.x)
						break
					}
				}
			}
		}
		return
	}
	var t [nMax][nMax]int
	for i := 0; i < nMax; i++ {
		for j := 0; j < nMax; j++ {
			t[i][j] = dis[i][j]
		}
	}
	mid := (l + r) >> 1
	for i := mid + 1; i <= r; i++ {
		Ins(i)
	}
	solve(l, mid)
	for i := 0; i < nMax; i++ {
		for j := 0; j < nMax; j++ {
			dis[i][j] = t[i][j]
		}
	}
	for i := l; i <= mid; i++ {
		Ins(i)
	}
	solve(mid+1, r)
	dis, t = t, dis
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 0; i < nMax; i++ {
		for j := 0; j < nMax; j++ {
			dis[i][j] = INF
		}
	}
	for i := 1; i <= m; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		dis[v][u] = w
		dis[u][v] = w
		g[u] = append(g[u], pair{w, v})
		g[v] = append(g[v], pair{w, u})
		deg[u]++
		deg[v]++
	}
	for i := 1; i <= n; i++ {
		sortPair(g[i])
	}
	solve(1, n)
	if res == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type pair struct {
	x, y int
}

func (l pair) lessThan(r pair) bool {
	if l.x == r.x {
		return l.y < r.y
	}
	return l.x < r.x
}

func (l pair) eq(r pair) bool {
	return l.x == r.x && l.y == r.y
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
