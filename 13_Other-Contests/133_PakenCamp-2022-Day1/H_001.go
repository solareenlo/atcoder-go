package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	A := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	q := &HeapPair{}
	for i := 0; i < N; i++ {
		heap.Push(q, pair{A[i], i})
	}
	ofs := 0
	for i := 0; i < Q; i++ {
		var t, x, y int
		fmt.Fscan(in, &t, &x, &y)
		x--
		if t == 1 {
			A[x] += y
			ofs += y
			if v[x] == 0 {
				heap.Push(q, pair{A[x], x})
			}
		} else {
			A[x] -= y
			if v[x] == 0 {
				heap.Push(q, pair{A[x], x})
			}
		}
		ans := make([]int, 0)
		for q.Len() > 0 {
			p := (*q)[0]
			if p.x != A[p.y] {
				heap.Pop(q)
				continue
			}
			if p.x > ofs {
				break
			}
			heap.Pop(q)
			if v[p.y] != 0 {
				continue
			}
			ans = append(ans, p.y)
			v[p.y] = 1
		}
		sort.Ints(ans)
		sz := len(ans)
		fmt.Fprint(out, sz)
		for j := 0; j < sz; j++ {
			fmt.Fprintf(out, " %d", ans[j]+1)
		}
		fmt.Fprintln(out)
	}
}

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int { return len(h) }
func (h HeapPair) Less(i, j int) bool {
	if h[i].x == h[j].x {
		return h[i].y < h[j].y
	}
	return h[i].x < h[j].x
}
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
