package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var par = [200200]int{}

func root(x int) int {
	if x == par[x] {
		return x
	}
	par[x] = root(par[x])
	return par[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	que := &Heap{}
	for i := 0; i < m; i++ {
		var ia, ib, ic int
		fmt.Fscan(in, &ia, &ib, &ic)
		heap.Push(que, edge{ia, ib, ic})
		heap.Push(que, edge{(ia + 1) % n, ib, ic + 1})
	}
	for i := 0; i < n; i++ {
		par[i] = i
	}
	ret := 0
	for que.Len() > 0 {
		e := (*que)[0]
		heap.Pop(que)
		if root(e.a) != root(e.b) {
			par[root(e.a)] = root(e.b)
			ret += e.c
			heap.Push(que, edge{(e.a + 1) % n, (e.b + 1) % n, e.c + 2})
		}
	}
	fmt.Println(ret)
}

type edge struct{ a, b, c int }
type Heap []edge

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].c < h[j].c }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(edge)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
