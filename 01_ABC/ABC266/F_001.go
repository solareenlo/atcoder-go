package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 400005

var cnt int
var vis [N]int
var v [][]int
var d [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	v = make([][]int, N)
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		v[a] = append(v[a], b)
		v[b] = append(v[b], a)
		d[a]++
		d[b]++
	}

	q := &Heap{}
	heap.Init(q)
	for i := 1; i <= n; i++ {
		if d[i] == 1 {
			heap.Push(q, i)
		}
	}
	for q.Len() > 0 {
		u := heap.Pop(q).(int)
		for _, x := range v[u] {
			d[x]--
			if d[x] == 1 {
				heap.Push(q, x)
			}
		}
	}

	for i := 1; i <= n; i++ {
		if d[i] == 2 {
			cnt++
			dfs(i, 0)
		}
	}

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		var a, b int
		fmt.Fscan(in, &a, &b)
		if vis[a] == vis[b] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
		t--
	}
}

func dfs(u, fa int) {
	vis[u] = cnt
	for _, x := range v[u] {
		if d[x] == 2 || x == fa {
			continue
		}
		dfs(x, u)
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
