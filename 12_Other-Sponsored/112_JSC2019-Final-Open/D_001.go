package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct {
	x, y int
}

func above(p, q, r P) bool {
	a := r.y - p.y
	b := -r.x + p.x
	c := p.y*r.x - p.x*r.y
	return a*q.x+b*q.y+c <= 0
}

func check(s, t P) bool {
	return s.y*t.x < s.x*t.y
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const SIZE = 200005
	const INF = 1000000000000000

	var pos [SIZE]P

	var n int
	fmt.Fscan(in, &n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	B := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &B[i])
	}
	sz := 0
	ret := -INF
	pos[sz] = P{0, B[0]}
	sz++
	for i := 1; i < n; i++ {
		l, r := 0, sz
		for r-l > 10 {
			a := (2*l + r) / 3
			b := (l + 2*r) / 3
			s := P{i - pos[a].x, A[i] - pos[a].y}
			t := P{i - pos[b].x, A[i] - pos[b].y}
			if check(s, t) {
				l = a
			} else {
				r = b
			}
		}
		for j := l; j < r; j++ {
			s := P{i - pos[j].x, A[i] - pos[j].y}
			if s.y >= 0 {
				ret = max(ret, (s.y+s.x-1)/s.x)
			} else {
				ret = max(ret, -((-s.y) / s.x))
			}
		}
		fmt.Fprintln(out, ret)
		p := P{i, B[i]}
		for sz >= 2 && above(pos[sz-2], pos[sz-1], p) {
			sz--
		}
		pos[sz] = p
		sz++
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
