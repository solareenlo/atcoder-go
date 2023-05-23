package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	ps := make([]point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ps[i].x, &ps[i].y, &ps[i].z)
	}
	cv := convex(ps)

	var up, lw Q
	up.init()
	lw.init()

	for _, tmp := range cv {
		a := tmp.x
		b := tmp.y
		c := tmp.z
		w := sign1(crs3(a, b, c).z)
		a = a.mul(k)
		b = b.mul(k)
		c = c.mul(k)
		if w > 0 {
			up.add(a, b, c)
		} else if w < 0 {
			lw.add(negz(a), negz(c), negz(b))
		}
	}
	up.calc()
	lw.calc()
	fmt.Println(plusMod(up.ans, lw.ans))
}

const eps = 0

func sign1(a int) int {
	if a < -eps {
		return -1
	}
	if a > eps {
		return 1
	}
	return 0
}

func sign2(a, b int) int { return sign1(a - b) }

type point struct {
	x, y, z int
}

func (l point) cmp(r point) bool {
	return l.x == r.x && l.y == r.y && l.z == r.z
}

func (l point) minus(r point) point {
	return point{l.x - r.x, l.y - r.y, l.z - r.z}
}

func (l point) div(v int) point {
	return point{l.x / v, l.y / v, l.z / v}
}

func (l point) lessThan(r point) bool {
	return 4*sign2(l.x, r.x)+2*sign2(l.y, r.y)+sign2(l.z, r.z) < 0
}

func (l point) mul(r int) point {
	return point{l.x * r, l.y * r, l.z * r}
}

type T struct {
	x, y, z point
}

type pairP struct {
	x, y point
}

func convex(ps []point) []T {
	ls := make([]point, 0)
	for _, p := range ps {
		if len(ls) <= 1 {
			ls = append(ls, p)
		} else if len(ls) == 2 {
			if !is0(crs3(ls[0], ls[1], p)) {
				ls = append(ls, p)
			}
		} else if len(ls) == 3 {
			if sign1(det4(ls[0], ls[1], ls[2], p)) != 0 {
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
		for _, tmp := range res {
			a := tmp.x
			b := tmp.y
			c := tmp.z
			if sign1(det4(p, a, b, c)) < 0 {
				del = append(del, pairP{a, b})
				del = append(del, pairP{b, c})
				del = append(del, pairP{c, a})
			} else {
				nx = append(nx, T{a, b, c})
			}
		}
		SortPair(del)
		for _, tmp := range del {
			a := tmp.x
			b := tmp.y
			idx := binarySearch(del, pairP{b, a})
			if idx == -1 {
				nx = append(nx, T{a, b, p})
			}
		}
		res, nx = nx, res
	}
	return res
}

func binarySearch(slice []pairP, value pairP) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2

		if slice[mid].x.cmp(value.x) && slice[mid].y.cmp(value.y) {
			return mid // 見つかった場合、インデックスを返す
		} else if slice[mid].x.lessThan(value.x) || (slice[mid].x.cmp(value.x) && slice[mid].y.lessThan(value.y)) {
			low = mid + 1 // 探索範囲を後半に絞る
		} else {
			high = mid - 1 // 探索範囲を前半に絞る
		}
	}

	return -1 // 見つからなかった場合、-1を返す
}

func SortPair(tmp []pairP) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y.lessThan(tmp[j].y)
		}
		return tmp[i].x.lessThan(tmp[j].x)
	})
}

type P struct {
	x pairP
	y int
}

type Q struct {
	ans int
	we  map[pairP]int
	wv  map[point]int
}

func (q *Q) init() {
	q.we = make(map[pairP]int)
	q.wv = make(map[point]int)
}

func (q *Q) add(a, b, c point) {
	q.ans = (q.ans + (work(a, b, c)+MOD)%MOD) % MOD
	q.we[minmax(a, b)] = (q.we[minmax(a, b)] + 1) % MOD
	q.we[minmax(b, c)] = (q.we[minmax(b, c)] + 1) % MOD
	q.we[minmax(c, a)] = (q.we[minmax(c, a)] + 1) % MOD
	q.wv[a] = (q.wv[a] + 1) % MOD
	q.wv[b] = (q.wv[b] + 1) % MOD
	q.wv[c] = (q.wv[c] + 1) % MOD
}

func work(a, b, c point) int {
	b = b.minus(a)
	c = c.minus(a)
	g := Gcd(Gcd(Gcd(b.x, b.y), Gcd(c.x, c.y)), Gcd(b.z, c.z))
	if g < 0 {
		g *= -1
	}
	b = b.div(g)
	c = c.div(g)
	d := b.x*c.y - b.y*c.x
	lx := minSlice(0, b.x, c.x) * 2
	rx := maxSlice(0, b.x, c.x) * 2
	ly := minSlice(0, b.y, c.y) * 2
	ry := maxSlice(0, b.y, c.y) * 2

	ns := 0
	ss := 0
	nt := 0
	st := 0

	for x := lx; x < rx+1; x++ {
		for y := ly; y < ry+1; y++ {
			p := x*c.y + y*(-c.x)
			q := x*(-b.y) + y*b.x
			z := fdiv(b.z*p+c.z*q, d)
			if 0 <= p && p < d && 0 <= q && q < d {
				ns = plusMod(ns, 1)
				ss = plusMod(ss, z)
			}
			if 0 <= p && 0 <= q && p+q <= d && q < d {
				nt = plusMod(nt, 1)
				st = plusMod(st, z)
			}
		}
	}

	res := val(minusMod(g, 1), plusMod(ss, mulMod(ns, a.z)), mulMod(ns, b.z), mulMod(ns, c.z))
	tmp0 := plusMod(st, mulMod(nt, a.z))
	res = plusMod(res, mulMod(tmp0, g))
	res = plusMod(res, DivMod(mulMod(mulMod(mulMod(nt, b.z), g), (g-1)), 2))
	res = plusMod(res, DivMod(mulMod(mulMod(mulMod(nt, c.z), g), (g-1)), 2))
	res = plusMod(res, plusMod(a.z, mulMod(c.z, g)))
	return res
}

func intMod(n int) int {
	if n < 0 {
		return (n + MOD) % MOD
	}
	return n % MOD
}

func mulMod(a, b int) int {
	return intMod(a) * intMod(b) % MOD
}

func plusMod(a, b int) int {
	return (intMod(a) + intMod(b)) % MOD
}

func minusMod(a, b int) int {
	return (intMod(a) - intMod(b) + MOD) % MOD
}

func val(k, s, x, y int) int {
	tmp0 := mulMod(DivMod(mulMod(k, plusMod(k, 1)), 2), s)
	tmp1 := mulMod(DivMod(mulMod(mulMod(plusMod(k, 1), k), minusMod(k, 1)), 6), plusMod(x, y))
	return plusMod(tmp0, tmp1)
}

func (q *Q) calc() {
	keys := make([]pairP, 0, len(q.we))
	for k := range q.we {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].x.cmp(keys[j].x) {
			return keys[i].y.lessThan(keys[j].y)
		}
		return keys[i].x.lessThan(keys[j].x)
	})
	for _, ab := range keys {
		a := ab.x
		b := ab.y
		val, _ := q.we[ab]
		val = minusMod(1, val)
		tmp := f(a, b)
		q.ans = plusMod(q.ans, mulMod(tmp, val))
		q.wv[a] = plusMod(q.wv[a], val)
		q.wv[b] = plusMod(q.wv[b], val)
	}
	keys1 := make([]point, 0, len(q.wv))
	for k := range q.wv {
		keys1 = append(keys1, k)
	}
	sort.Slice(keys1, func(i, j int) bool {
		return keys1[i].lessThan(keys1[j])
	})
	for _, p := range keys1 {
		val := q.wv[p]
		val = minusMod(1, val)
		q.ans = plusMod(q.ans, mulMod(p.z, val))
	}
}

func f(a, b point) int {
	b = b.minus(a)
	g := Gcd(Gcd(b.x, b.y), b.z)
	if g < 0 {
		g = -g
	}
	b = b.div(g)
	h := Gcd(b.x, b.y)
	if h < 0 {
		h = -h
	}

	n := 0
	s := 0
	for i := 0; i < h; i++ {
		n = plusMod(n, 1)
		s = plusMod(s, fdiv(b.z*i, h))
	}

	res := mulMod(plusMod(s, mulMod(n, a.z)), g)
	res = plusMod(res, DivMod(mulMod(mulMod(mulMod(n, b.z), g), minusMod(g, 1)), 2))
	res = plusMod(res, plusMod(a.z, mulMod(b.z, g)))
	return res
}

func fdiv(a, b int) int {
	if (a^b) < 0 && a%b != 0 {
		return a/b - 1
	}
	return a / b
}

func minmax(a, b point) pairP {
	if a.lessThan(b) {
		return pairP{a, b}
	}
	return pairP{b, a}
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func minSlice(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func maxSlice(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

func crs2(a, b point) point {
	return point{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
	}
}

func crs3(a, b, c point) point {
	return crs2(b.minus(a), c.minus(a))
}

func det3(a, b, c point) int {
	return dot(crs2(a, b), c)
}

func det4(a, b, c, d point) int {
	return det3(b.minus(a), c.minus(a), d.minus(a))
}

func dot(a, b point) int {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func is0(a point) bool {
	return sign1(a.x) == 0 && sign1(a.y) == 0 && sign1(a.z) == 0
}

func negz(a point) point {
	return point{a.x, a.y, 1 - a.z}
}

func DivMod(a, b int) int {
	ret := a * ModInv(b)
	ret %= MOD
	return ret
}

func ModInv(a int) int {
	b, u, v := MOD, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}
