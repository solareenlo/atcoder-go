package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)

	type pair struct{ x, y int }
	st := map[pair]bool{}

	x, y := 0, 0
	for _, c := range s {
		st[pair{x, y}] = true
		switch c {
		case 'L':
			x--
		case 'R':
			x++
		case 'D':
			y--
		case 'U':
			y++
		}
	}
	st[pair{x, y}] = true

	if x == 0 && y == 0 {
		fmt.Println(len(st))
		return
	}
	ps := make([]pair, 0, len(st))
	for k := range st {
		ps = append(ps, k)
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].x < ps[j].x || (ps[i].x == ps[j].x && ps[i].y < ps[j].y)
	})

	if x == 0 {
		x, y = y, x
		for i := range ps {
			ps[i].x, ps[i].y = ps[i].y, ps[i].x
		}
	}
	if x < 0 {
		x = -x
		for i := range ps {
			ps[i].x *= -1
		}
	}
	if y < 0 {
		y = -y
		for i := range ps {
			ps[i].y *= -1
		}
	}

	m := map[pair][]pair{}
	for _, p := range ps {
		nx, ny := p.x, p.y
		if nx > 0 {
			i := nx / x
			nx -= x * i
			ny -= y * i
		} else {
			i := (-nx + x - 1) / x
			nx += x * i
			ny += y * i
		}
		m[pair{nx, ny}] = append(m[pair{nx, ny}], p)
	}

	res := 0
	for _, np := range m {
		d := make([]int, 0)
		for _, p := range np {
			d = append(d, (p.x+len(s))/x)
		}
		sort.Ints(d)
		for i := 0; i < len(d); i++ {
			w := k
			if i+1 < len(d) {
				w = min(w, d[i+1]-d[i])
			}
			res += w
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
