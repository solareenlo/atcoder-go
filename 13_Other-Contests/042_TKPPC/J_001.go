package main

import (
	"bufio"
	"fmt"
	"os"
)

var inf = int(1e18)

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer OUT.Flush()

	var q int
	fmt.Fscan(IN, &q)
	parent := make([]int, 1)
	parent[0] = -1
	qry := make([][3]int, q)

	for i := 0; i < q; i++ {
		var o, a, c int
		fmt.Fscan(IN, &o, &a, &c)
		qry[i] = [3]int{o, a, c}
		if o == 1 {
			parent = append(parent, a)
		}
	}

	n := len(parent)
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		adj[parent[i]] = append(adj[parent[i]], i)
	}
	eul := make([]int, 0)
	in := make([]int, n)
	out := make([]int, n)
	pos := make([][]int, n)
	var dfs func(int)
	dfs = func(x int) {
		in[x] = len(eul)
		eul = append(eul, x)
		out[x] = in[x]
		pos[x] = append(pos[x], in[x])
		for _, y := range adj[x] {
			dfs(y)
			out[x] = len(eul)
			pos[x] = append(pos[x], out[x])
			eul = append(eul, x)
		}
	}
	dfs(0)

	seg := NewLazySegmentTree(2*n - 1)
	for _, x := range pos[0] {
		seg.modify(1, 0, seg.n, x, NewInfo(0))
	}

	cost := make([]int, n)
	cnt := 0
	for i := range qry {
		o := qry[i][0]
		a := qry[i][1]
		c := qry[i][2]
		if o == 1 {
			d := seg.rangeQuery(1, 0, seg.n, in[a], in[a]+1)
			dep := d.sid + c
			cnt++
			for _, x := range pos[cnt] {
				seg.modify(1, 0, seg.n, x, NewInfo(dep))
			}
			cost[cnt] = c
		} else if o == 2 {
			seg.rangeApply(1, 0, seg.n, in[a], out[a]+1, NewTag(c-cost[a]))
			cost[a] = c
		} else {
			dep := seg.rangeQuery(1, 0, seg.n, in[a], in[a]+1).sid
			tmp1 := seg.rangeQuery(1, 0, seg.n, 0, out[a]+1).l
			tmp2 := seg.rangeQuery(1, 0, seg.n, in[a], 2*n-1).r
			ans := max(tmp1, tmp2) + dep
			fmt.Fprintln(OUT, ans)
		}
	}
}

type LazySegmentTree struct {
	n    int
	info []INFO
	tag  []Tag
}

func NewLazySegmentTree(n int) *LazySegmentTree {
	L := new(LazySegmentTree)
	L.n = n
	L.info = make([]INFO, 4<<lg(n))
	for i := range L.info {
		L.info[i].sid = -inf
		L.info[i].mid = -inf
		L.info[i].l = -inf
		L.info[i].r = -inf
	}
	L.tag = make([]Tag, 4<<lg(n))
	return L
}

func (L *LazySegmentTree) pull(p int) {
	L.info[p] = Merge(L.info[2*p], L.info[2*p+1])
}

func (L *LazySegmentTree) apply(p int, v Tag) {
	L.info[p].apply(v)
	L.tag[p].apply(v)
}

func (L *LazySegmentTree) push(p int) {
	L.apply(2*p, L.tag[p])
	L.apply(2*p+1, L.tag[p])
	L.tag[p] = NewTag(0)
}

func (L *LazySegmentTree) modify(p, l, r, x int, v INFO) {
	if r-l == 1 {
		L.info[p] = v
		return
	}
	m := (l + r) / 2
	L.push(p)
	if x < m {
		L.modify(2*p, l, m, x, v)
	} else {
		L.modify(2*p+1, m, r, x, v)
	}
	L.pull(p)
}

func (L *LazySegmentTree) rangeQuery(p, l, r, x, y int) INFO {
	if l >= y || r <= x {
		return NewInfoInf()
	}
	if l >= x && r <= y {
		return L.info[p]
	}
	m := (l + r) / 2
	L.push(p)
	return Merge(L.rangeQuery(2*p, l, m, x, y), L.rangeQuery(2*p+1, m, r, x, y))
}

func (L *LazySegmentTree) rangeApply(p, l, r, x, y int, v Tag) {
	if l >= y || r <= x {
		return
	}
	if l >= x && r <= y {
		L.apply(p, v)
		return
	}
	m := (l + r) / 2
	L.push(p)
	L.rangeApply(2*p, l, m, x, y, v)
	L.rangeApply(2*p+1, m, r, x, y, v)
	L.pull(p)
}

type Tag struct {
	add int
}

func NewTag(x int) Tag {
	return Tag{x}
}

func (T *Tag) apply(t Tag) {
	T.add += t.add
}

type INFO struct {
	sid, mid, l, r int
}

func NewInfoInf() INFO {
	return INFO{sid: -inf, mid: -inf, l: -inf, r: -inf}
}

func NewInfo(d int) INFO {
	return INFO{sid: d, mid: -2 * d, l: -d, r: -d}
}

func (I *INFO) apply(t Tag) {
	I.sid += t.add
	I.mid -= 2 * t.add
	I.l -= t.add
	I.r -= t.add
}

func Merge(a, b INFO) INFO {
	var c INFO
	c.sid = max(a.sid, b.sid)
	c.mid = max(a.mid, b.mid)
	c.l = max(a.l, b.l, a.sid+b.mid)
	c.r = max(a.r, b.r, a.mid+b.sid)
	return c
}

func lg(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
