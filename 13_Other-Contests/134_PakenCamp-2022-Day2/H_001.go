package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf = 1 << 55

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	A := make([]int, N)
	B := make([]int, N-1)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	for i := range B {
		fmt.Fscan(in, &B[i])
	}
	var Q int
	fmt.Fscan(in, &Q)
	T := make([]int, Q)
	P := make([]int, Q)
	L := make([]int, Q)
	R := make([]int, Q)
	V := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &T[i], &P[i], &V[i], &L[i], &R[i])
		P[i]--
		L[i]--
	}
	res := solve(N, A, B, Q, T, P, V, L, R)
	for _, e := range res {
		fmt.Println(e + 1)
	}
}

func solve(N int, A, B []int, Q int, T, P, V, L, R []int) []int {
	initVec := make([]Node, 0)
	bId := make(map[int]int)
	initVec = append(initVec, Node{inf, -inf})
	for i := 0; i < N; i++ {
		initVec = append(initVec, Node{A[i], inf - 1})
		if i != N-1 {
			initVec = append(initVec, Node{0, B[i]})
		}
		if i != N-1 {
			bId[B[i]] = i
		}
	}
	initVec = append(initVec, Node{inf, -inf})
	seg := NewSegTreeNode(initVec, eNode, opNode)

	var advance func(*int, *int)
	advance = func(l, r *int) {
		mn := seg.Prod(2*(*l)+1, 2*(*r)).Bmin
		i := bId[mn]
		sumL := seg.Prod(2*(*l)+1, 2*i+2).Asum
		sumR := seg.Prod(2*i+3, 2*(*r)).Asum
		if sumL >= sumR {
			*r = i + 1
		} else {
			*l = i + 1
		}
	}

	var skip func(*int, *int)
	skip = func(l, r *int) {
		Asum := seg.Prod(2*(*l)+1, 2*(*r)).Asum
		skt := (Asum + 2) / 2
		m := seg.MaxRight(2*(*l)+1, func(x Node) bool { return x.Asum < skt }).x
		tmp := seg.specialSearch(m, skt)
		nl := tmp.x
		nr := tmp.y
		nv := tmp.z
		mn := inf - 1
		if nv.Bmin != inf {
			mn = nv.Bmin
		}
		var f func(Node) bool
		f = func(e Node) bool {
			return e.Bmin > mn
		}
		pl := nl
		if nl/2 != (*l) {
			pl = seg.MinLeft(nl, f).x
		}
		pr := nr
		if (nr-2)/2 != (*r) {
			pr = seg.MaxRight(nr, f).x
		}
		*l = pl / 2
		*r = pr / 2
	}

	var calcAns func(int, int) int
	calcAns = func(l, r int) int {
		for l+1 != r {
			skip(&l, &r)
			if l+1 != r {
				advance(&l, &r)
			}
		}
		return l
	}

	ans := make([]int, 0)
	for q := 0; q < Q; q++ {
		if T[q] == 1 {
			A[P[q]] = V[q]
			seg.Set(2*P[q]+1, Node{V[q], inf - 1})
		}
		if T[q] == 2 {
			delete(bId, B[P[q]])
			B[P[q]] = V[q]
			bId[V[q]] = P[q]
			seg.Set(2*P[q]+2, Node{0, V[q]})
		}
		lv := -inf
		if L[q] != 0 {
			lv = B[L[q]-1]
		}
		rv := -inf
		if R[q] != N {
			rv = B[R[q]-1]
		}
		seg.Set(2*L[q], Node{inf, -inf})
		seg.Set(2*R[q], Node{inf, -inf})
		ans = append(ans, calcAns(L[q], R[q]))
		if L[q] == 0 {
			seg.Set(2*L[q], Node{inf, lv})
		} else {
			seg.Set(2*L[q], Node{0, lv})
		}
		if R[q] == N {
			seg.Set(2*R[q], Node{inf, rv})
		} else {
			seg.Set(2*R[q], Node{0, rv})
		}
	}
	return ans
}

type Node struct {
	Asum, Bmin int
}

func opNode(a, b Node) Node {
	return Node{a.Asum + b.Asum, min(a.Bmin, b.Bmin)}
}

func eNode() Node { return Node{0, inf} }

type ENode func() Node
type OpNode func(a, b Node) Node
type CompareNode func(v Node) bool
type SegTreeNode struct {
	n    int
	size int
	log  int
	d    []Node
	e    ENode
	op   OpNode
}

func NewSegTreeNode(n []Node, e ENode, op OpNode) *SegTreeNode {
	seg := new(SegTreeNode)
	seg.n = len(n)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]Node, 2*seg.size)
	seg.e = e
	seg.op = op
	for i, _ := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		seg.d[seg.size+i] = n[i]
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.Update(i)
	}
	return seg
}

func (seg *SegTreeNode) Update(k int) {
	seg.d[k] = seg.op(seg.d[2*k], seg.d[2*k+1])
}

func (seg *SegTreeNode) Set(p int, x Node) {
	if p < 0 || seg.n <= p {
		panic("")
	}
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}

func (seg *SegTreeNode) Get(p int) Node {
	if p < 0 || seg.n <= p {
		panic("")
	}
	return seg.d[p+seg.size]
}

func (seg *SegTreeNode) Prod(l, r int) Node {
	if l < 0 || r < l || seg.n < r {
		panic("")
	}
	sml, smr := seg.e(), seg.e()
	l += seg.size
	r += seg.size
	for l < r {
		if (l & 1) == 1 {
			sml = seg.op(sml, seg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = seg.op(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return seg.op(sml, smr)
}

func (seg *SegTreeNode) AllProd() Node {
	return seg.d[1]
}

type pairNode struct {
	x int
	y Node
}

func (seg *SegTreeNode) MaxRight(l int, cmp CompareNode) pairNode {
	if l < 0 || seg.n < l {
		panic("")
	}
	if !cmp(seg.e()) {
		panic("")
	}
	if l == seg.n {
		return pairNode{seg.n, seg.e()}
	}
	l += seg.size
	sm := seg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(seg.op(sm, seg.d[l])) {
			for l < seg.size {
				l = 2 * l
				if cmp(seg.op(sm, seg.d[l])) {
					sm = seg.op(sm, seg.d[l])
					l++
				}
			}
			return pairNode{l - seg.size, sm}
		}
		sm = seg.op(sm, seg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return pairNode{seg.n, sm}
}

func (seg *SegTreeNode) MinLeft(r int, cmp CompareNode) pairNode {
	if r < 0 || seg.n < r {
		panic("")
	}
	if !cmp(seg.e()) {
		panic("")
	}
	if r == 0 {
		return pairNode{0, seg.e()}
	}
	r += seg.size
	sm := seg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(seg.op(seg.d[r], sm)) {
			for r < seg.size {
				r = 2*r + 1
				if cmp(seg.op(seg.d[r], sm)) {
					sm = seg.op(seg.d[r], sm)
					r--
				}
			}
			return pairNode{r + 1 - seg.size, sm}
		}
		sm = seg.op(seg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return pairNode{0, sm}
}

func (seg *SegTreeNode) ceilPow2(n int) int {
	x := 0
	for (uint(1) << x) < uint(n) {
		x++
	}
	return x
}

type tuple struct {
	x, y int
	z    Node
}

func (seg *SegTreeNode) specialSearch(m, v int) tuple {
	pr := seg.e()
	l := m + seg.size
	r := m + seg.size
	l--
	lup := true
	rup := true
	for r%2 == 0 {
		r /= 2
	}
	for l > 1 && l%2 == 1 {
		l /= 2
	}
	for {
		npr := seg.op(seg.d[l], seg.op(pr, seg.d[r]))
		if npr.Asum < v {
			if seg.d[l].Bmin > seg.d[r].Bmin {
				pr = seg.op(seg.d[l], pr)
				if lup {
					l--
					for l > 1 && l%2 == 1 {
						l /= 2
					}
				} else if l < seg.size {
					l = 2 * l
				} else {
					break
				}
			} else {
				pr = seg.op(pr, seg.d[r])
				if rup {
					r++
					for r%2 == 0 {
						r /= 2
					}
				} else if r < seg.size {
					r = 2*r + 1
				} else {
					break
				}
			}
		} else {
			if seg.d[l].Bmin < seg.d[r].Bmin {
				if lup {
					lup = false
				}
				if l < seg.size {
					l = 2*l + 1
				} else {
					l++
					break
				}
			} else {
				if rup {
					rup = false
				}
				if r < seg.size {
					r = 2 * r
				} else {
					break
				}
			}
		}
	}
	if l >= seg.size && !lup {
		l -= seg.size
		tmp := seg.MaxRight(l, func(e Node) bool { return e.Asum < v })
		r = tmp.x
		pr = tmp.y
	} else {
		r -= seg.size
		tmp := seg.MinLeft(r, func(e Node) bool { return e.Asum < v })
		l = tmp.x
		pr = tmp.y
	}
	if l == r {
		for pr.Asum < v {
			vl := -1
			if l != 0 {
				vl = seg.d[l-1+seg.size].Bmin
			}
			vr := -1
			if r != seg.n {
				vr = seg.d[r+seg.size].Bmin
			}
			if vl > vr {
				l--
				pr = seg.op(seg.d[l+seg.size], pr)
			} else {
				pr = seg.op(pr, seg.d[r+seg.size])
				r++
			}
		}
		return tuple{l, r, pr}
	}
	if seg.d[l+seg.size].Asum == 0 {
		l++
	}
	if seg.d[r-1+seg.size].Asum == 0 {
		r--
	}
	pr = seg.Prod(l, r)
	for pr.Asum < v {
		vl := -1
		if l != 0 {
			vl = seg.d[l-1+seg.size].Bmin
		}
		vr := -1
		if r != seg.n {
			vr = seg.d[r+seg.size].Bmin
		}
		if vl > vr {
			l--
			pr = seg.op(seg.d[l+seg.size], pr)
		} else {
			pr = seg.op(pr, seg.d[r+seg.size])
			r++
		}
	}
	return tuple{l, r, pr}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
