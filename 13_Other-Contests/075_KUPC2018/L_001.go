package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	ls := make([]Point, 0)
	var n int
	fmt.Fscan(in, &n)
	ps := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ps[i].x, &ps[i].y, &ps[i].z)
	}
	cv := convex(ps)
	tmp := make(map[Point]Point)
	for _, t := range cv {
		a, b, c := t.x, t.y, t.z
		w := crs3(a, b, c)
		tmp[normdir(w)] = tmp[normdir(w)].plus(w)
	}
	for _, val := range tmp {
		ls = append(ls, val)
	}
	n = len(ls)
	cand := make(map[Point]bool)
	for i := 0; i < n; i++ {
		cand[normdirabs(ls[i])] = true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			d := normdirabs(crs2(ls[i], ls[j]))
			if !is0(d) {
				cand[d] = true
			}
		}
	}
	mx := -1
	bits := make(map[Point]bool)
	var work func(bit []bool)
	work = func(bit []bool) {
		var sum Point
		for i := 0; i < n; i++ {
			if bit[i] {
				sum = sum.plus(ls[i])
			}
		}
		w := norm(sum)
		if mx < w {
			mx = w
			bits = make(map[Point]bool)
		}
		if mx == w {
			bits[normdirabs(sum)] = true
		}
	}
	for v, _ := range cand {
		base := make([]bool, n)
		idx := make([]int, 0)
		for i := 0; i < n; i++ {
			w := dot(v, ls[i])
			if w > 0 {
				base[i] = true
			} else if w == 0 {
				idx = append(idx, i)
			}
		}
		work(base)
		for _, i := range idx {
			now := make([]int, 0)
			for _, j := range idx {
				if det3(ls[i], ls[j], v) > 0 {
					now = append(now, j)
				}
			}
			sort.Slice(now, func(j, k int) bool {
				return det3(ls[now[j]], ls[now[k]], v) > 0
			})
			cur := make([]bool, len(base))
			copy(cur, base)
			cur[i] = true
			work(cur)
			for _, j := range now {
				cur[j] = true
				work(cur)
			}
		}
	}
	fmt.Println(math.Sqrt(float64(mx)) / 2.0)
	fmt.Println(len(bits) * 2)
}

func normdir(a Point) Point {
	if is0(a) {
		return a
	}
	g := gcd(gcd(a.x, a.y), a.z)
	return a.div(g)
}

func normdirabs(a Point) Point {
	a = normdir(a)
	if a.lessThan(Point{-a.x, -a.y, -a.z}) {
		return Point{-a.x, -a.y, -a.z}
	}
	return a
}

type Point struct {
	x, y, z int
}

func (p Point) plus(r Point) Point  { return Point{p.x + r.x, p.y + r.y, p.z + r.z} }
func (p Point) minus(r Point) Point { return Point{p.x - r.x, p.y - r.y, p.z - r.z} }
func (p Point) div(v int) Point     { return Point{p.x / v, p.y / v, p.z / v} }

func (p Point) lessThan(r Point) bool {
	return 4*sgn2(p.x, r.x)+2*sgn2(p.y, r.y)+sgn2(p.z, r.z) < 0
}

func sgn1(c int) int {
	if c < 0 {
		return -1
	}
	if c > 0 {
		return 1
	}
	return 0
}

func sgn2(a, b int) int {
	return sgn1(a - b)
}

func dot(a, b Point) int {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func norm(a Point) int {
	return dot(a, a)
}

func is0(a Point) bool {
	return sgn1(a.x) == 0 && sgn1(a.y) == 0 && sgn1(a.z) == 0
}

func crs2(a, b Point) Point {
	return Point{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
	}
}

func crs3(a, b, c Point) Point {
	return crs2(b.minus(a), c.minus(a))
}

func det3(a, b, c Point) int {
	return dot(crs2(a, b), c)
}

func det4(a, b, c, d Point) int {
	return det3(b.minus(a), c.minus(a), d.minus(a))
}

type T struct {
	x, y, z Point
}

func convex(ps []Point) []T {
	ls := make([]Point, 0)
	for _, p := range ps {
		if len(ls) <= 1 {
			ls = append(ls, p)
		} else if len(ls) == 2 {
			if !is0(crs3(ls[0], ls[1], p)) {
				ls = append(ls, p)
			}
		} else if len(ls) == 3 {
			if sgn1(det4(ls[0], ls[1], ls[2], p)) != 0 {
				ls = append(ls, p)
			}
		}
	}
	res := make([]T, 0)
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 4; k++ {
				a := ls[i]
				b := ls[j]
				c := ls[k]
				d := ls[6-i-j-k]
				if det4(d, a, b, c) < 0 {
					a, b = b, a
				}
				res = append(res, T{a, b, c})
			}
		}
	}
	for _, p := range ps {
		nx := make([]T, 0)
		del := make([]pairP, 0)
		for _, t := range res {
			a, b, c := t.x, t.y, t.z
			if sgn1(det4(p, a, b, c)) < 0 {
				del = append(del, pairP{a, b})
				del = append(del, pairP{b, c})
				del = append(del, pairP{c, a})
			} else {
				nx = append(nx, T{a, b, c})
			}
		}
		sort.Slice(del, func(i, j int) bool {
			if reflect.DeepEqual(del[i].x, del[j].x) {
				return del[i].y.lessThan(del[j].y)
			}
			return del[i].x.lessThan(del[j].x)
		})
		for _, t := range del {
			a, b := t.x, t.y
			idx := sort.Search(len(del), func(i int) bool {
				return !del[i].lessThan(pairP{b, a})
			})
			if idx < len(del) && reflect.DeepEqual(del[idx].x, b) && reflect.DeepEqual(del[idx].y, a) {
			} else {
				nx = append(nx, T{a, b, p})
			}
		}
		res = nx
	}
	return res
}

type pairP struct {
	x, y Point
}

func (a pairP) lessThan(b pairP) bool {
	if reflect.DeepEqual(a.x, b.x) {
		return a.y.lessThan(b.y)
	}
	return a.x.lessThan(b.x)
}

func gcd(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return gcd(b, a%b)
}
