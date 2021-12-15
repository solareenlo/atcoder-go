package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	ou := bufio.NewWriter(os.Stdout)
	defer ou.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	indeg := make([]int, n)
	out := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a -= 1
		b -= 1
		indeg[b] += 1
		out[a] = append(out[a], b)
	}

	pq := &Heap{}
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			heap.Push(pq, i)
		}
	}

	res := make([]int, 0)
	for pq.Len() > 0 {
		i := (*pq)[0]
		heap.Pop(pq)
		res = append(res, i)
		for _, j := range out[i] {
			indeg[j] -= 1
			if indeg[j] == 0 {
				heap.Push(pq, j)
			}
		}
	}

	if len(res) != n {
		fmt.Fprintln(ou, -1)
	} else {
		for i := 0; i < n; i++ {
			fmt.Fprint(ou, res[i]+1, " ")
		}
		fmt.Fprintln(ou)
	}
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
