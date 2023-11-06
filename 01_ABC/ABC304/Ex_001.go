package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 200200

var vis [N]bool
var r [N]int
var g [N][]int

func dfs(u int) {
	vis[u] = true
	for _, v := range g[u] {
		if !vis[v] {
			dfs(v)
		}
		r[u] = min(r[u], r[v]-1)
	}
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var l, in, ans [N]int
	var vec [N][]int

	var n, m int
	fmt.Fscan(IN, &n, &m)
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(IN, &x, &y)
		g[x] = append(g[x], y)
		in[y]++
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(IN, &l[i], &r[i])
	}
	for i := 1; i <= n; i++ {
		if !vis[i] {
			dfs(i)
		}
	}
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			vec[l[i]] = append(vec[l[i]], i)
		}
	}
	q := make(HeapPair, 0)
	heap.Init(&q)
	for i := 1; i <= n; i++ {
		for _, x := range vec[i] {
			heap.Push(&q, pair{r[x], x})
		}
		if q.Len() == 0 {
			fmt.Fprintln(out, "No")
			return
		}
		u := heap.Pop(&q).(pair).y
		if r[u] < i {
			fmt.Fprintln(out, "No")
			return
		}
		ans[u] = i
		for _, v := range g[u] {
			in[v]--
			if in[v] == 0 {
				vec[max(l[v], i+1)] = append(vec[max(l[v], i+1)], v)
			}
		}
	}
	fmt.Fprintln(out, "Yes")
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", ans[i])
	}
	fmt.Fprintln(out)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
