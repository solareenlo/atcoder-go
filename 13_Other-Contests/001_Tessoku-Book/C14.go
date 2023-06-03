package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]pair, n)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 << 30
	}
	used := make([]bool, n)
	pq := &HeapPair{}
	nxt := make([]int, 0)
	ans := 0
	for m > 0 {
		m--
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		u--
		v--
		g[u] = append(g[u], pair{v, c})
		g[v] = append(g[v], pair{u, c})
	}
	dp[0] = 0
	heap.Push(pq, pair{0, 0})
	for pq.Len() > 0 {
		tmp := heap.Pop(pq).(pair)
		d := tmp.x
		v := tmp.y
		if d > dp[v] {
			continue
		}
		for _, p := range g[v] {
			u := p.x
			c := p.y
			if d+c >= dp[u] {
				continue
			}
			dp[u] = d + c
			heap.Push(pq, pair{dp[u], u})
		}
	}
	nxt = append(nxt, n-1)
	used[n-1] = true
	for len(nxt) != 0 {
		v := nxt[0]
		nxt = nxt[1:]
		ans++
		for _, p := range g[v] {
			u := p.x
			c := p.y
			if used[u] || dp[v] != dp[u]+c {
				continue
			}
			nxt = append(nxt, u)
			used[u] = true
		}
	}
	fmt.Println(ans)
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
