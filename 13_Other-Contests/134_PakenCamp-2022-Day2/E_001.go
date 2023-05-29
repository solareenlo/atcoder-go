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

	var N, M int
	fmt.Fscan(in, &N, &M)

	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &B[i], &C[i])
		B[i]--
	}

	ord := make([]int, N)
	for i := range ord {
		ord[i] = i
	}
	sort.Slice(ord, func(i, j int) bool {
		return A[ord[i]] < A[ord[j]]
	})

	f := make([]int, N-M+1)
	max := make([]int, M)
	h := &Heap{}
	sum := 0
	for i := N - 1; i >= 0; i-- {
		j := ord[i]
		x := C[j]
		if x > max[B[j]] {
			x, max[B[j]] = max[B[j]], x
		}
		heap.Push(h, x)
		sum += x
		if h.Len() > M {
			sum -= heap.Pop(h).(int)
		}
		if h.Len() == M {
			f[i] = sum
		}
	}

	var Q int
	fmt.Fscan(in, &Q)

	for Q > 0 {
		Q--
		var x int
		fmt.Fscan(in, &x)
		it := upperBound(f, x)
		if it < len(f) && f[it] == f[0] {
			fmt.Println(0)
		} else {
			fmt.Println(A[ord[it-1]])
		}
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] > h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
