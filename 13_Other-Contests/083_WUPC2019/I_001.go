package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const U = 1 << 17

type point struct {
	x, y int
}

var S [U * 2][]point
var O, C, D, X, P [100000]int
var E [U][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &O[i], &C[i], &D[i], &X[i])
		if C[i] == -1 {
			C[i] = U - 2
		}
		E[O[i]] = append(E[O[i]], i)
	}
	meguru(1, 0, U)
	for i := 0; i < N; i++ {
		fmt.Println(P[i])
	}
}

func meguru(k, l, r int) {
	S[k] = convex_hull(S[k])
	if r-l == 1 {
		for _, i := range E[l] {
			mn := X[i]
			for j := 0; j <= 17; j++ {
				mn = min(mn, query(S[k>>j], -2*D[i])+D[i]*D[i])
			}
			P[i] = mn
			insert(l+1, C[i]+2, point{D[i], P[i] + D[i]*D[i]})
		}
		return
	}
	m := (l + r) / 2
	meguru(k*2+0, l, m)
	meguru(k*2+1, m, r)
}

func convex_hull(ps []point) []point {
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].x == ps[j].x {
			return ps[i].y > ps[j].y
		}
		return ps[i].x > ps[j].x
	})
	hs := make([]point, 0)
	for _, p := range ps {
		for len(hs) >= 2 && cross(minus(hs[len(hs)-1], hs[len(hs)-2]), minus(p, hs[len(hs)-1])) >= 0 {
			hs = hs[:len(hs)-1]
		}
		hs = append(hs, p)
	}
	return hs
}

func cross(a, b point) int {
	return a.x*b.y - a.y*b.x
}

func minus(a, b point) point {
	return point{a.x - b.x, a.y - b.y}
}

func insert(l, r int, a point) {
	l += U
	r += U - 1
	for l <= r {
		if (l & 1) == 1 {
			S[l] = append(S[l], a)
			l++
		}
		if (r & 1) == 0 {
			S[r] = append(S[r], a)
			r--
		}
		l >>= 1
		r >>= 1
	}
}

func query(ps []point, x int) int {
	if len(ps) == 0 {
		return 1e18
	}
	l := -1
	r := len(ps) - 1
	for r-l > 1 {
		m := (l + r) / 2
		if f(ps[m], x) <= f(ps[m+1], x) {
			r = m
		} else {
			l = m
		}
	}
	return f(ps[r], x)
}

func f(a point, x int) int {
	return a.x*x + a.y
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
