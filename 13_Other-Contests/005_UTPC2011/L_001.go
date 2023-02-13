package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005
const inf = 2147483647

var n, q, K, idx int
var a, GO, nxt, head, sz, dep, root, f, son, dfn, l, top, r []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &q)
	a = make([]int, N*2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	GO = make([]int, N*2)
	nxt = make([]int, N*2)
	head = make([]int, N)
	sz = make([]int, N)
	dep = make([]int, N)
	root = make([]int, N)
	f = make([]int, N)
	son = make([]int, N)
	dfn = make([]int, N)
	l = make([]int, N)
	r = make([]int, N)
	top = make([]int, N)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		add(u, v)
		add(v, u)
	}
	dfs1(1, 0)
	dfs2(1, 1)
	lans := 0
	for q > 0 {
		q--
		var u, v, k int
		fmt.Fscan(in, &u, &v, &k)
		p := lca(u, v)
		lans = T.query(root[u], root[v], root[p], root[f[p]], 1, inf, k)
		fmt.Fprintln(out, lans)
	}
}

func dfs1(x, fa int) {
	sz[x] = 1
	dep[x] = dep[fa] + 1
	root[x] = root[fa]
	f[x] = fa
	T.modify(&root[x], 1, inf, a[x], 1)
	maxn := pair{0, 0}
	for i := head[x]; i > 0; i = nxt[i] {
		g := GO[i]
		if g == fa {
			continue
		}
		dfs1(g, x)
		sz[x] += sz[g]
		if sz[g] > maxn.x || (sz[g] == maxn.x && g > maxn.x) {
			maxn = pair{sz[g], g}
		}
	}
	son[x] = maxn.y
}

func dfs2(x, fa int) {
	if x == 0 {
		return
	}
	idx++
	dfn[idx] = x
	l[x] = idx
	top[x] = fa
	dfs2(son[x], fa)
	for i := head[x]; i > 0; i = nxt[i] {
		g := GO[i]
		if g != f[x] && g != son[x] {
			dfs2(g, g)
		}
	}
	r[x] = idx
}

func lca(x, y int) int {
	for top[x] != top[y] {
		if dep[top[x]] < dep[top[y]] {
			x, y = y, x
		}
		x = f[top[x]]
	}
	if dep[x] < dep[y] {
		return x
	}
	return y
}

func add(u, v int) {
	K++
	nxt[K] = head[u]
	head[u] = K
	GO[K] = v
}

type pair struct {
	x, y int
}

var T segTree

type node struct {
	lc, rc, val int
}

type segTree struct {
	tree [50000005]node
	cnt  int
}

func (s *segTree) pushup(p int) {
	s.tree[p].val = s.tree[s.tree[p].lc].val + s.tree[s.tree[p].rc].val

}

func (s *segTree) clone(p *int) {
	s.cnt++
	s.tree[s.cnt] = s.tree[*p]
	*p = s.cnt
}

func (s *segTree) modify(p *int, l, r, x, k int) {
	s.clone(p)
	if l == r {
		s.tree[*p].val += k
		return
	}
	mid := (l + (r-l)/2)
	if x <= mid {
		s.modify(&s.tree[*p].lc, l, mid, x, k)
	} else {
		s.modify(&s.tree[*p].rc, mid+1, r, x, k)
	}
	s.pushup(*p)
}

func (s *segTree) query(p1, p2, p3, p4, l, r, k int) int {
	sz := s.tree[s.tree[p1].lc].val + s.tree[s.tree[p2].lc].val - s.tree[s.tree[p3].lc].val - s.tree[s.tree[p4].lc].val
	if l == r {
		return l
	}
	mid := (l + (r-l)/2)
	if sz >= k {
		return s.query(s.tree[p1].lc, s.tree[p2].lc, s.tree[p3].lc, s.tree[p4].lc, l, mid, k)
	} else {
		return s.query(s.tree[p1].rc, s.tree[p2].rc, s.tree[p3].rc, s.tree[p4].rc, mid+1, r, k-sz)
	}
}
