package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var r, b int
	fmt.Fscan(in, &r, &b)
	w := r + b
	var s string
	fmt.Fscan(in, &s)
	sr := make([]int, n+1)
	sb := make([]int, n+1)
	for i := 0; i < n; i++ {
		sr[i+1] = sr[i]
		sb[i+1] = sb[i]
		if i%w == 0 {
			sr[i+1] += r
		}
		if i%w == 0 {
			sb[i+1] += b
		}
		if s[i] == 'R' {
			sr[i+1]--
		}
		if s[i] == 'B' {
			sb[i+1]--
		}
	}
	dp := make([]int, n+1)
	d := make([]bit2, w)
	for i := 0; i < n; i++ {
		d[i%w].reg(sr[i], sb[i])
	}
	for i := 0; i < w; i++ {
		d[i].init()
	}
	for i := 0; i < n; i++ {
		d[i%w].add(sr[i], sb[i], dp[i]-i/w+n)
		ni := i + 1
		dp[ni] = d[ni%w].sum(sr[ni], sb[ni]) + ni/w - n
		dp[ni] = max(dp[ni], dp[i])
	}
	fmt.Println(dp[n])
}

type Pair struct {
	x, y int
}

type bit2 struct {
	n  int
	d  []bit
	xs X
	ys []X
	p  []Pair
}

func (b2 *bit2) reg(x, y int) { b2.p = append(b2.p, Pair{x, y}) }

func (b2 *bit2) init() {
	for i := 0; i < len(b2.p); i++ {
		b2.xs.add(b2.p[i].x)
	}
	b2.xs.init()
	b2.n = b2.xs.size() + 1
	resizeX(&b2.ys, b2.n)
	for j := 0; j < len(b2.p); j++ {
		for i := b2.xs.op(b2.p[j].x); i < b2.n; i += i & -i {
			b2.ys[i].add(b2.p[j].y)
		}
	}
	resizeBit(&b2.d, b2.n)
	for i := 0; i < b2.n; i++ {
		b2.ys[i].init()
		var tmp bit
		tmp.init(b2.ys[i].size() + 2)
		b2.d[i] = tmp
	}
}

func (b2 *bit2) add(x, y, z int) {
	for i := b2.xs.op(x); i < b2.n; i += i & -i {
		b2.d[i].add(b2.ys[i].op(y), z)
	}
}

func (b2 *bit2) sum(x, y int) int {
	res := 0
	for i := b2.xs.op(x); i > 0; i -= i & -i {
		res = max(res, b2.d[i].sum(b2.ys[i].op(y)))
	}
	return res
}

type bit struct {
	n int
	d []int
}

func (b *bit) init(mx int) {
	b.n = mx
	b.d = make([]int, mx)
}

func (b *bit) add(i, x int) {
	for i = i + 1; i < b.n; i += i & -i {
		b.d[i] = max(b.d[i], x)
	}
}

func (b *bit) sum(i int) int {
	x := 0
	for i = i + 1; i > 0; i -= i & -i {
		x = max(x, b.d[i])
	}
	return x
}

type X struct {
	d []int
}

func (xx *X) add(x int) { xx.d = append(xx.d, x) }

func (x *X) init() {
	sort.Ints(x.d)
	x.d = unique(x.d)
}

func (x X) size() int { return len(x.d) }

func (xx X) op(x int) int { // !! upper bound (1-indexed) !!
	return UpperBound(xx.d, x)
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

func UpperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func resizeX(a *[]X, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			var tmp X
			tmp.d = make([]int, 0)
			*a = append(*a, tmp)
		}
	}
}

func resizeBit(a *[]bit, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			var tmp bit
			tmp.init(0)
			*a = append(*a, tmp)
		}
	}
}
