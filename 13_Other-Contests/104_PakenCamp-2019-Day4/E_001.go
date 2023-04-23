package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

const INF = int(1e18)
const MOD = 1_000_000_007

type st struct {
	u int
	c [2]int
}

var n, m int
var E [300000][]st
var d [2][300000]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		a--
		b--
		E[a] = append(E[a], st{b, [2]int{c, d}})
		E[b] = append(E[b], st{a, [2]int{c, d}})
	}
	for i := range d {
		for j := range d[i] {
			d[i][j] = INF
		}
	}
	dijkstra(0)
	dijkstra(1)
	fmt.Println((calc(0, 1) + calc(1, 0) + MOD - calc(0, 0)) % MOD)
}

func dijkstra(b int) {
	que := &HeapPair{}
	d[b][0] = 0
	heap.Push(que, pair{0, 0})
	for que.Len() > 0 {
		p := heap.Pop(que).(pair)
		if d[b][p.y] != p.x {
			continue
		}
		for _, u := range E[p.y] {
			if d[b][u.u] > p.x+u.c[b] {
				d[b][u.u] = p.x + u.c[b]
				heap.Push(que, pair{d[b][u.u], u.u})
			}
		}
	}
}

func calc(A, B int) int {
	var G [300000][]int
	for i := 0; i < n; i++ {
		for _, u := range E[i] {
			if (A != 0 || d[0][u.u] == d[0][i]-u.c[0]) && (B != 0 || d[1][u.u] == d[1][i]-u.c[1]) {
				G[u.u] = append(G[u.u], i)
			}
		}
	}
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = i
	}
	sort.Slice(v, func(a, b int) bool {
		if A == 0 {
			return d[0][v[a]] < d[0][v[b]]
		}
		return d[1][v[a]] < d[1][v[b]]
	})
	var dp [300000]int
	dp[0] = 1
	for i := 0; i < n; i++ {
		for _, u := range G[v[i]] {
			dp[u] = (dp[u] + dp[v[i]]) % MOD
		}
	}
	return dp[n-1]
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
