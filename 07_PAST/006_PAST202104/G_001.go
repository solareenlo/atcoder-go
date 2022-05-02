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

	var N, M, Q int
	fmt.Fscan(in, &N, &M, &Q)

	adj := make([][]pair, 200005)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		adj[a] = append(adj[a], pair{b, c})
		adj[b] = append(adj[b], pair{a, c})
	}

	q := &Heap{}
	add := make([]pair, 0)

	vis := make([]int, 200005)
	vis[1] = 1
	for i := range adj[1] {
		b := adj[1][i].x
		c := adj[1][i].y
		heap.Push(q, pair{-c, b})
	}

	cnt := 1
	for i := 0; i < Q; i++ {
		var x int
		fmt.Fscan(in, &x)
		for q.Len() > 0 {
			d := -(*q)[0].x
			a := (*q)[0].y
			if d > x {
				break
			}
			heap.Pop(q)
			if vis[a] != 0 {
				continue
			}
			vis[a] = 1
			cnt++
			for i := range adj[a] {
				b := adj[a][i].x
				c := adj[a][i].y
				add = append(add, pair{-c, b})
			}
		}
		for len(add) > 0 {
			heap.Push(q, add[0])
			add = add[1:]
		}
		fmt.Fprintln(out, cnt)
	}
}

type pair struct{ x, y int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
