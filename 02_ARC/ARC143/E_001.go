package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 200005

var G, F [][]int
var col, in []int

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer OUT.Flush()

	var n int
	fmt.Fscan(IN, &n)

	G = make([][]int, N)
	for i := 2; i <= n; i++ {
		var u, v int
		fmt.Fscan(IN, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	var S string
	fmt.Fscan(IN, &S)
	S = "_" + S
	col = make([]int, N)
	for i := 1; i <= n; i++ {
		if S[i] == 'W' {
			col[i] = 1
		} else {
			col[i] = 0
		}
	}

	F = make([][]int, N)
	in = make([]int, N)
	dfs(1, 0)
	if col[1] == 0 {
		fmt.Fprintln(OUT, -1)
		return
	}

	q := &Heap{}
	heap.Init(q)
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			heap.Push(q, -i)
		}
	}
	for q.Len() > 0 {
		u := -heap.Pop(q).(int)
		fmt.Fprintf(OUT, "%d ", u)
		for i := 0; i < len(F[u]); i++ {
			v := F[u][i]
			in[v]--
			if in[v] == 0 {
				heap.Push(q, -v)
			}
		}
	}
}

func dfs(u, fa int) {
	for i := 0; i < len(G[u]); i++ {
		v := G[u][i]
		if v == fa {
			continue
		}
		dfs(v, u)
	}
	if fa != 0 {
		if col[u] != 0 {
			col[fa] ^= 1
			F[u] = append(F[u], fa)
			in[fa]++
		} else {
			F[fa] = append(F[fa], u)
			in[u]++
		}
	}
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
