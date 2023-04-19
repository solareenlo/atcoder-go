package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	X := make([]int, N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i], &P[i])
	}

	var Q int
	fmt.Fscan(in, &Q)
	var T UFTree
	T.init(N, X, P)
	for i := 0; i < Q; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		T.unite(a, b)
		fmt.Println(T.ans)
	}
}

type UFTree struct {
	_n            int
	pq_ng         []*HeapNG
	pq_ok         []*HeapOK
	out, wei, par []int
	component     int
	ans           int
}

func (uf *UFTree) init(n int, X, P []int) {
	uf._n = n
	uf.pq_ng = make([]*HeapNG, n)
	for i := 0; i < n; i++ {
		uf.pq_ng[i] = &HeapNG{}
	}
	uf.pq_ok = make([]*HeapOK, n)
	for i := 0; i < n; i++ {
		uf.pq_ok[i] = &HeapOK{}
	}
	uf.out = make([]int, n)
	uf.wei = make([]int, n)
	uf.par = make([]int, n)
	uf.component = n
	for i := 0; i < n; i++ {
		uf.wei[i] = 1
		uf.par[i] = i
		if X[i] >= 0 {
			uf.out[i] = X[i]
			uf.wei[i] = 0
		} else {
			heap.Push(uf.pq_ng[i], pair{P[i], -X[i]})
		}
	}
	uf.ans = 0
}

func (uf *UFTree) root(a int) int {
	if a == uf.par[a] {
		return a
	}
	uf.par[a] = uf.root(uf.par[a])
	return uf.par[a]
}

func (uf *UFTree) same(a, b int) int {
	if uf.root(a) == uf.root(b) {
		return 1
	}
	return 0
}

func (uf *UFTree) unite(a, b int) bool {
	a = uf.root(a)
	b = uf.root(b)
	if a == b {
		return false
	}
	if uf.wei[a] < uf.wei[b] {
		a, b = b, a
	}
	uf.par[b] = a

	uf.out[a] += uf.out[b]
	for uf.pq_ng[b].Len() > 0 {
		heap.Push(uf.pq_ng[a], heap.Pop(uf.pq_ng[b]).(pair))
	}
	for uf.pq_ok[b].Len() > 0 {
		heap.Push(uf.pq_ok[a], heap.Pop(uf.pq_ok[b]).(pair))
	}
	for uf.pq_ng[a].Len() > 0 && uf.pq_ok[a].Len() > 0 {
		if (*uf.pq_ng[a])[0].x <= (*uf.pq_ok[a])[0].x {
			break
		}
		n_ok := heap.Pop(uf.pq_ng[a]).(pair)
		n_ng := heap.Pop(uf.pq_ok[a]).(pair)
		var tmp pair
		v := min(n_ng.y, n_ok.y)
		uf.ans += (n_ok.x - n_ng.x) * v
		tmp = pair{n_ok.x, v}
		for uf.pq_ok[a].Len() > 0 && (*uf.pq_ok[a])[0].x == n_ok.x {
			tmp.y += heap.Pop(uf.pq_ok[a]).(pair).y
		}
		heap.Push(uf.pq_ok[a], tmp)
		if v != n_ok.y {
			tmp.y = n_ok.y - v
			heap.Push(uf.pq_ng[a], tmp)
		}
		tmp = pair{n_ng.x, v}
		for uf.pq_ng[a].Len() > 0 && (*uf.pq_ng[a])[0].x == n_ng.x {
			tmp.y += heap.Pop(uf.pq_ng[a]).(pair).y
		}
		heap.Push(uf.pq_ng[a], tmp)
		if v != n_ng.y {
			tmp.y = n_ng.y - v
			heap.Push(uf.pq_ok[a], tmp)
		}
	}
	for uf.out[a] != 0 && uf.pq_ng[a].Len() > 0 {
		tmp := heap.Pop(uf.pq_ng[a]).(pair)
		if uf.out[a] >= tmp.y {
			uf.out[a] -= tmp.y
			uf.ans += tmp.y * tmp.x
			heap.Push(uf.pq_ok[a], tmp)
		} else {
			heap.Push(uf.pq_ok[a], pair{tmp.x, uf.out[a]})
			heap.Push(uf.pq_ng[a], pair{tmp.x, tmp.y - uf.out[a]})
			uf.ans += tmp.x * uf.out[a]
			uf.out[a] = 0
		}
	}
	uf.wei[a] += uf.wei[b]
	uf.component--
	return true
}

type pair struct {
	x, y int
}

type HeapNG []pair

func (h HeapNG) Len() int            { return len(h) }
func (h HeapNG) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapNG) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapNG) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapNG) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type HeapOK []pair

func (h HeapOK) Len() int            { return len(h) }
func (h HeapOK) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapOK) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapOK) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapOK) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
