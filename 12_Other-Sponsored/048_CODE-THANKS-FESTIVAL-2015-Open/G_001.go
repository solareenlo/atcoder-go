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

	const MX = 40000

	var A, B, C, T [MX * 2]int
	var sum [MX]int

	var N, M int
	fmt.Fscan(in, &N, &M)

	cs := make([][]int, MX)
	cs[0] = append(cs[0], 1)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &A[i], &B[i], &C[i], &T[i])
		A[i]--
		B[i]--
		cs[A[i]] = append(cs[A[i]], C[i])
		cs[B[i]] = append(cs[B[i]], C[i])
	}

	sz := 0
	G := make([][]pair, 3<<17)
	for i := 0; i < N; i++ {
		sum[i] = sz
		sort.Ints(cs[i])
		cs[i] = unique(cs[i])
		for j := 1; j < len(cs[i]); j++ {
			G[sz+j] = append(G[sz+j], pair{sz + j - 1, cs[i][j] - cs[i][j-1]})
			G[sz+j-1] = append(G[sz+j-1], pair{sz + j, cs[i][j] - cs[i][j-1]})
		}
		sz += len(cs[i])
	}
	for i := 0; i < M; i++ {
		u := lowerBound(cs[A[i]], C[i]) + sum[A[i]]
		v := lowerBound(cs[B[i]], C[i]) + sum[B[i]]
		G[u] = append(G[u], pair{v, T[i]})
		G[v] = append(G[v], pair{u, T[i]})
	}
	var d [3 << 17]int
	for i := 1; i < sz; i++ {
		d[i] = 9e18
	}
	P := make(HeapPair, 1)
	P[0] = pair{0, 0}
	heap.Init(&P)
	for P.Len() > 0 {
		tmp := heap.Pop(&P).(pair)
		c := -tmp.x
		u := tmp.y
		if d[u] < c {
			continue
		}
		if u >= sum[N-1] {
			fmt.Println(c)
			return
		}
		for _, e := range G[u] {
			v := e.x
			if d[v] > c+e.y {
				d[v] = c + e.y
				heap.Push(&P, pair{-d[v], v})
			}
		}
	}
}

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
