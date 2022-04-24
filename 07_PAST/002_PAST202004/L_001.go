package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, K, D int
	fmt.Fscan(in, &N, &K, &D)

	que := &Heap{}
	A := make([]int, N)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	if D*(K-1)+1 > N {
		fmt.Fprintln(out, -1)
		return
	}

	p, l := 0, 0
	for n := K; n >= 1; n-- {
		for i := p; i <= (N-1)-(n-1)*D; i++ {
			heap.Push(que, pair{-A[i], -i})
		}
		p = (N - 1) - (n-1)*D + 1
		for -(*que)[0].y < l {
			heap.Pop(que)
		}
		l = -(*que)[0].y + D
		fmt.Fprint(out, -(*que)[0].x, " ")
		heap.Pop(que)
	}
}

type pair struct{ x, y int }

type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x || (h[i].x == h[j].x && h[i].y > h[j].y) }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
