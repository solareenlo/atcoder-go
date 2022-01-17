package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const N = 1000005
const M = 200005

var (
	cnt int
	rt  int
	p   int
	a   = [N]int{}
	s   = [N]int{}
	b   = [M]int{}
	t   = [N]int{}
	lc  = [N]int{}
	rc  = [N]int{}
	sz  = [N]int{}
)

func newnode(x int) int {
	cnt++
	u := cnt
	sz[u] = 1
	a[u] = x
	s[u] = x
	lc[u] = 0
	rc[u] = 0
	t[u] = 0
	return u
}

func copy(u int) int {
	cnt++
	v := cnt
	a[v] = a[u]
	s[v] = s[u]
	lc[v] = lc[u]
	rc[v] = rc[u]
	sz[v] = sz[u]
	t[v] = t[u]
	return v
}

func pushup(u int) {
	sz[u] = sz[lc[u]] + sz[rc[u]] + 1
	s[u] = s[lc[u]] + s[rc[u]] + a[u]
}

func add(u, k int) {
	a[u] += k
	t[u] += k
	s[u] += k * sz[u]
}

func pushdown(u int) {
	if t[u] == 0 {
		return
	}
	if lc[u] != 0 {
		lc[u] = copy(lc[u])
		add(lc[u], t[u])
	}
	if rc[u] != 0 {
		rc[u] = copy(rc[u])
		add(rc[u], t[u])
	}
	t[u] = 0
}

func merge(u, v int) int {
	if u == 0 || v == 0 {
		return u | v
	}
	if rand.Intn(1<<60)%(sz[u]+sz[v]) < sz[u] {
		pushdown(u)
		u = copy(u)
		rc[u] = merge(rc[u], v)
		pushup(u)
		return u
	} else {
		pushdown(v)
		v = copy(v)
		lc[v] = merge(u, lc[v])
		pushup(v)
		return v
	}
}

func split(u, k int, a, b *int) {
	if u == 0 {
		*a = 0
		*b = 0
		return
	}
	pushdown(u)
	if sz[lc[u]] < k {
		*a = copy(u)
		split(rc[*a], k-sz[lc[*a]]-1, &rc[*a], b)
		pushup(*a)
	} else {
		*b = copy(u)
		split(lc[*b], k, a, &lc[*b])
		pushup(*b)
	}
}

func upd(l, r, k int) {
	var x, y, z int
	split(rt, l-1, &x, &y)
	split(y, r-l+1, &y, &z)
	y = copy(y)
	add(y, k)
	rt = merge(merge(x, y), z)
}

func pas(l1, r1, l2, r2 int) {
	if l1 == l2 && r1 == r2 {
		return
	}
	var x1, y1, z1, x2, y2, z2 int
	split(rt, l1-1, &x1, &y1)
	split(y1, r1-l1+1, &y1, &z1)
	split(rt, l2-1, &x2, &y2)
	split(y2, r2-l2+1, &y2, &z2)
	rt = merge(merge(x2, y1), z2)
}

func query(l, r int) int {
	var x, y, z int
	var ans int
	split(rt, l-1, &x, &y)
	split(y, r-l+1, &y, &z)
	ans = s[y]
	return ans
}

func dfs(u int) {
	pushdown(u)
	if lc[u] != 0 {
		dfs(lc[u])
	}
	p++
	b[p] = a[u]
	if rc[u] != 0 {
		dfs(rc[u])
	}
}

func build(l, r int) int {
	if l > r {
		return 0
	}
	mid := (l + r) >> 1
	u := newnode(b[mid])
	lc[u] = build(l, mid-1)
	rc[u] = build(mid+1, r)
	pushup(u)
	return u
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &b[i])
	}
	rt = build(1, n)
	for i := 0; i < q; i++ {
		var ty int
		fmt.Fscan(in, &ty)
		if ty == 1 {
			var l, r, k int
			fmt.Fscan(in, &l, &r, &k)
			upd(l, r, k)
		}
		if ty == 2 {
			var l1, r1, l2, r2 int
			fmt.Fscan(in, &l1, &r1, &l2, &r2)
			pas(l2, r2, l1, r1)
		}
		if ty == 3 {
			var l, r int
			fmt.Fscan(in, &l, &r)
			fmt.Fprintln(out, query(l, r))
		}
		if cnt > 700000 {
			p = 0
			dfs(rt)
			rt = 0
			cnt = 0
			rt = build(1, n)
		}
	}
}
