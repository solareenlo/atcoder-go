package main

import (
	"bufio"
	"fmt"
	"os"
)

const MX = 200000
const MAX_N = 1 << 18
const INF = int(1e18)

type Node struct {
	a, b int
}

func op(l, r Node) Node {
	return Node{l.a + r.a, min(l.b+r.a, r.b)}
}

type SegTree1 struct {
	lazy [MAX_N*2 - 1]Node
}

func (seg *SegTree1) Init() {
	for i := range seg.lazy {
		seg.lazy[i] = Node{0, INF}
	}
}

func (seg *SegTree1) setLazy(k int, x Node) {
	seg.lazy[k] = op(seg.lazy[k], x)
}

func (seg *SegTree1) push(k int) {
	if seg.lazy[k].b == INF {
		return
	}
	seg.setLazy(k*2+1, seg.lazy[k])
	seg.setLazy(k*2+2, seg.lazy[k])
	seg.lazy[k] = Node{0, INF}
}

func (seg *SegTree1) add(a, b int, x Node, k, l, r int) {
	if b <= l || r <= a {
		return
	}
	if a <= l && r <= b {
		seg.setLazy(k, x)
		return
	}
	seg.push(k)
	seg.add(a, b, x, k*2+1, l, (l+r)/2)
	seg.add(a, b, x, k*2+2, (l+r)/2, r)
}

type SegTree2 struct {
	seg  [MAX_N*2 - 1]int
	lazy [MAX_N*2 - 1]int
}

func (seg *SegTree2) setLazy(k, x int) {
	seg.lazy[k] += x
	seg.seg[k] += x
}

func (seg *SegTree2) push(k int) {
	if seg.lazy[k] == 0 {
		return
	}
	seg.setLazy(k*2+1, seg.lazy[k])
	seg.setLazy(k*2+2, seg.lazy[k])
	seg.lazy[k] = 0
}

func (seg *SegTree2) add(a, b, x, k, l, r int) {
	if b <= l || r <= a {
		return
	}
	if a <= l && r <= b {
		seg.setLazy(k, x)
		return
	}
	seg.push(k)
	seg.add(a, b, x, k*2+1, l, (l+r)/2)
	seg.add(a, b, x, k*2+2, (l+r)/2, r)
	seg.seg[k] = max(seg.seg[k*2+1], seg.seg[k*2+2])
}

func (seg *SegTree2) query(a, b, k, l, r int) int {
	if b <= l || r <= a {
		return -INF
	}
	if a <= l && r <= b {
		return seg.seg[k]
	}
	seg.push(k)
	return max(seg.query(a, b, k*2+1, l, (l+r)/2), seg.query(a, b, k*2+2, (l+r)/2, r))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	var L, R, X, Y, A [MX]int
	var seg SegTree1
	seg.Init()
	var seg2 SegTree2
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &L[i], &R[i], &X[i], &Y[i])
		L[i]--
		R[i]--
		seg.add(L[i], R[i]+1, Node{X[i], Y[i]}, 0, 0, MAX_N)
	}
	for k := 0; k < MAX_N-1; k++ {
		seg.push(k)
	}
	for i := 0; i < N; i++ {
		x := seg.lazy[i+MAX_N-1]
		A[i] = x.b - x.a
		seg2.add(i, i+1, A[i], 0, 0, MAX_N)
	}
	for i := 0; i < Q; i++ {
		seg2.add(L[i], R[i]+1, X[i], 0, 0, MAX_N)
		if seg2.query(L[i], R[i]+1, 0, 0, MAX_N) != Y[i] {
			fmt.Fprintln(out, "NG")
			return
		}
	}
	fmt.Fprintln(out, "OK")
	for i := 0; i < N; i++ {
		fmt.Fprint(out, A[i], " ")
	}
	fmt.Fprintln(out)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
