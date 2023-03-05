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

	const INF = int(1e18)
	const MAX = 30005

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	var s, g int
	fmt.Fscan(in, &s, &g)

	G := make([][]pair, MAX)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		G[a] = append(G[a], pair{b, c})
		G[b] = append(G[b], pair{a, c})
	}

	A := make([]int, N)
	dis := make([]int, MAX)
	for i := 0; i < N; i++ {
		dis[i] = INF
	}

	var dijkstra func(int)
	dijkstra = func(u int) {
		dis[u] = 0
		PQ := &HeapPair{}
		heap.Push(PQ, pair{0, u})
		for PQ.Len() > 0 {
			tmp := heap.Pop(PQ).(pair)
			a, b := tmp.x, tmp.y
			if dis[b] < a {
				continue
			}
			for i := 0; i < len(G[b]); i++ {
				c := G[b][i].x
				d := G[b][i].y
				if dis[c] > dis[b]+d {
					dis[c] = dis[b] + d
					heap.Push(PQ, pair{dis[c], c})
				}
			}
		}
		return
	}
	dijkstra(s)
	for i := 0; i < N; i++ {
		A[i] = dis[i]
		dis[i] = INF
	}
	dijkstra(g)

	S := make([]pair, K)
	for i := 0; i < K; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		S[i] = pair{a, b}
	}
	S = append(S, pair{INF, INF})

	ans := INF
	for m := 0; m < N; m++ {
		sum := 0
		it := lowerBound(S, pair{A[m] + 1, -1})
		it--
		if it >= 0 {
			sum += S[it].y
		}
		if m == g {
			ans = min(ans, sum)
		}
		it = lowerBound(S, pair{dis[m] + 1, -1})
		it--
		if it >= 0 {
			sum += S[it].y
		}
		ans = min(ans, sum)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x
	})
	return idx
}

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
