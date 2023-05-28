package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 1001001001

type P struct {
	x, y int
}

type Q struct {
	x, y P
}

func main() {
	in := bufio.NewReader(os.Stdin)

	initfact()

	var n, m int
	fmt.Fscan(in, &n, &m)
	p := make([]Q, m)
	for i := range p {
		fmt.Fscan(in, &p[i].x.x, &p[i].x.y, &p[i].y.x, &p[i].y.y)
		p[i].x.x--
		p[i].x.y--
		p[i].y.x--
		p[i].y.y--
	}

	var f func(P, P) int
	f = func(a, b P) int {
		if a.x == b.x && a.y == b.y {
			return 0
		}
		if a.x == b.x {
			return 1
		}
		if a.x == b.y {
			return 1
		}
		if a.y == b.x {
			return 1
		}
		if a.y == b.y {
			return 1
		}
		return 0
	}

	ts := NewDsu(n * 2)
	t := NewDsu(n * 2)
	for _, tmp := range p {
		ai := tmp.x.x
		aj := tmp.x.y
		bi := tmp.y.x
		bj := tmp.y.y
		bi += n
		bj += n
		t.Merge(ai, aj)
		t.Merge(bi, bj)
		ts.Merge(ai, aj)
		ts.Merge(bi, bj)
		ts.Merge(ai, bi)
	}
	g := make([]int, n*2)
	ms := make([]int, n*2)
	for i := range ms {
		ms[i] = INF
	}
	for a := 0; a < m; a++ {
		for b := 0; b < a; b++ {
			e1 := f(p[a].x, p[b].x)
			e2 := f(p[a].y, p[b].y)
			if e1 != e2 {
				g[ts.Leader(p[a].x.x)] = 1
			}
		}
	}
	es := make([][]Q, n*2)
	for _, tmp := range p {
		p1 := tmp.x
		p2 := tmp.y
		es[ts.Leader(p1.x)] = append(es[ts.Leader(p1.x)], Q{p1, P{p2.x + n, p2.y + n}})
	}

	ans := 1
	fs := make([]int, n*2)
	for i := range fs {
		fs[i] = 1
	}
	for i := 0; i < n*2; i++ {
		if t.Leader(i) == i {
			fs[ts.Leader(i)] = fs[ts.Leader(i)] * facs[t.Size(i)] % MOD
			if g[ts.Leader(i)] == 0 {
				ms[ts.Leader(i)] = min(ms[ts.Leader(i)], t.Size(i))
			}
		}
	}
	for i := 0; i < n*2; i++ {
		if ts.Leader(i) == i {
			if g[i] != 0 {
				ans = ans * (fs[i] * invs[2] % MOD) % MOD
			} else {
				if ts.Size(i) > 8 {
					ans = ans * (fs[i] * ifacs[ms[i]] % MOD) % MOD
				} else {
					ans = ans * solve(es[i]) % MOD
				}
			}
		}
	}
	fmt.Println(ans)
}

func solve(es []Q) int {
	var xs X
	for _, tmp := range es {
		p1 := tmp.x
		p2 := tmp.y
		xs.add(p1.x)
		xs.add(p1.y)
		xs.add(p2.x)
		xs.add(p2.y)
	}
	for i := range es {
		es[i].x.x = xs.ope2(es[i].x.x)
		es[i].x.y = xs.ope2(es[i].x.y)
		es[i].y.x = xs.ope2(es[i].y.x)
		es[i].y.y = xs.ope2(es[i].y.y)
	}

	st := make(map[[12]int]struct{})
	q := make([][12]int, 0)
	var push func([12]int)
	push = func(a [12]int) {
		if _, ok := st[a]; ok {
			return
		}
		st[a] = struct{}{}
		q = append(q, a)
	}
	push(pm())
	for len(q) > 0 {
		a := q[0]
		q = q[1:]
		for _, tmp := range es {
			ai := tmp.x.x
			aj := tmp.x.y
			bi := tmp.y.x
			bj := tmp.y.y
			a[ai], a[aj] = a[aj], a[ai]
			a[bi], a[bj] = a[bj], a[bi]
			push(a)
			a[ai], a[aj] = a[aj], a[ai]
			a[bi], a[bj] = a[bj], a[bi]
		}
	}
	return len(st)
}

func pm() [12]int {
	var a [12]int
	for i := range a {
		a[i] = i
	}
	return a
}

type X struct {
	ini bool
	d   []int
}

func (x *X) add(a int) {
	x.d = append(x.d, a)
}

func (x *X) init() {
	sort.Ints(x.d)
	x.d = unique(x.d)
	x.ini = true
}

func (x *X) size() int {
	if !x.ini {
		x.init()
	}
	return len(x.d)
}

func (x *X) ope1(i int) int {
	if !x.ini {
		x.init()
	}
	return x.d[i]
}

func (x *X) ope2(a int) int {
	if !x.ini {
		x.init()
	}
	return upperBound(x.d, a) - 1
}

const vmax = (1 << 21) + 10

var facs, ifacs, invs [vmax]int

func initfact() {
	facs[0] = 1
	for i := 1; i < vmax; i++ {
		facs[i] = facs[i-1] * i % MOD
	}
	ifacs[vmax-1] = InvMod(facs[vmax-1])
	for i := vmax - 2; i >= 0; i-- {
		ifacs[i] = ifacs[i+1] * (i + 1) % MOD
	}
	for i := vmax - 1; i >= 1; i-- {
		invs[i] = ifacs[i] * facs[i-1] % MOD
	}
}

type dsu struct {
	n            int
	parentOrSize []int
}

func NewDsu(n int) *dsu {
	d := new(dsu)
	d.n = n
	d.parentOrSize = make([]int, d.n)
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *dsu) Merge(a, b int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	x := d.Leader(a)
	y := d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d *dsu) Same(a, b int) bool {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	return d.Leader(a) == d.Leader(b)
}

func (d *dsu) Leader(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *dsu) Size(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return -d.parentOrSize[d.Leader(a)]
}

func (d *dsu) Groups() [][]int {
	leaderBuf := make([]int, d.n)
	groupSize := make([]int, d.n)
	for i := 0; i < d.n; i++ {
		leaderBuf[i] = d.Leader(i)
		groupSize[leaderBuf[i]]++
	}
	result := make([][]int, d.n)
	for i := 0; i < d.n; i++ {
		result[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < d.n; i++ {
		result[leaderBuf[i]] = append(result[leaderBuf[i]], i)
	}
	eraseEmpty := func(a [][]int) [][]int {
		result := make([][]int, 0, len(a))
		for i := range a {
			if len(a[i]) != 0 {
				result = append(result, a[i])
			}
		}
		return result
	}
	return eraseEmpty(result)
}

const MOD = 998244353

func PowMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func InvMod(a int) int {
	return PowMod(a, MOD-2)
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
