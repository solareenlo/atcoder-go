package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

const MX1 = 210000
const MX2 = 110000
const INF = 999999999

var g [MX1]pair
var fs, nx [MX1]int
var v [MX2]bool
var Q []pair
var center, fi, se, m, ans, at, k, X int
var dist []int
var num [MX2]int

func count(a, b int) int {
	sort.Ints(dist[a:b])
	b--
	ret, cur := 0, 0
	for i := a + 1; i <= b; i++ {
		cur += dist[i]
	}
	for a < b {
		if dist[a]+dist[b] <= k {
			ret -= k*(b-a) - dist[a]*(b-a) - cur
			cur -= dist[a+1]
			a++
		} else {
			cur -= dist[b]
			b--
		}
	}
	return ret
}

func Find(a, b, c int) {
	num[at] = a
	dist[at] = c
	at++
	last := at
	for i := fs[a]; i >= 0; i = nx[i] {
		if v[g[i].x] && b != g[i].x {
			Find(g[i].x, a, g[i].y+c)
			if b == -1 {
				Q = append(Q, pair{last, at})
				ans -= count(last, at)
				last = at
			}
		}
	}
}

func dfi(a, b int) int {
	ret, p := 1, 0
	for i := fs[a]; i >= 0; i = nx[i] {
		if v[g[i].x] && b != g[i].x {
			t := dfi(g[i].x, a)
			ret += t
			p = max(p, t)
		}
	}
	p = max(p, se-fi-p)
	if p < m {
		m = p
		center = a
	}
	return ret
}

var L, sz [MX2]int

func main() {
	in := bufio.NewReader(os.Stdin)

	dist = make([]int, MX2)

	var a, b int
	fmt.Fscan(in, &a, &b)
	for i := 0; i < a; i++ {
		fs[i] = -1
	}
	for i := 0; i < a; i++ {
		L[i] = INF
	}
	for i := 0; i < a*2; i++ {
		nx[i] = -1
	}
	k = b
	use := 0
	for i := 0; i < a-1; i++ {
		var c, d, e int
		fmt.Fscan(in, &c, &d, &e)
		sz[c-1]++
		sz[d-1]++
		L[c-1] = min(L[c-1], e)
		L[d-1] = min(L[d-1], e)
		g[use] = pair{d - 1, e}
		nx[use] = fs[c-1]
		fs[c-1] = use
		use++
		g[use] = pair{c - 1, e}
		nx[use] = fs[d-1]
		fs[d-1] = use
		use++
	}
	Q = append(Q, pair{0, a})
	for i := 0; i < a; i++ {
		num[i] = i
	}
	ans = k * a * (a - 1) / 2
	for len(Q) > 0 {
		fi = Q[0].x
		se = Q[0].y
		Q = Q[1:]
		if fi == se-1 {
			continue
		}
		for i := fi; i < se; i++ {
			v[num[i]] = true
		}
		m = INF
		dfi(num[fi], -1)
		at = fi
		Find(center, -1, 0)
		ans += count(fi, se)
		for i := fi; i < se; i++ {
			v[num[i]] = false
		}
	}
	av := 0
	for i := 0; i < a; i++ {
		for j := fs[i]; j >= 0; j = nx[j] {
			if g[j].y > b {
				best := g[j].y
				best = min(best, b+min(L[i], L[g[j].x]))
				if sz[i]+sz[g[j].x] != a {
					best = min(best, 2*b)
				}
				av += best - b
			}
		}
	}
	ans += av / 2
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
