package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var N, M int
	fmt.Fscan(in, &N, &M)

	type pair struct {
		x, y int
	}
	E := make([][]pair, N)
	for i := 0; i < M; i++ {
		var f, t, c int
		fmt.Fscan(in, &f, &t, &c)
		E[f] = append(E[f], pair{c, t})
		E[t] = append(E[t], pair{c, f})
	}

	d := make([][]int, 28)
	for i := range d {
		d[i] = make([]int, N)
		for j := range d[i] {
			d[i][j] = INF
		}
	}

	pq := &Heap{}
	heap.Init(pq)
	heap.Push(pq, tuple{0, 0, 0})
	for pq.Len() > 0 {
		tmp := heap.Pop(pq).(tuple)
		c := tmp.c
		t := tmp.t
		v := tmp.v
		if d[t][v] == INF {
			d[t][v] = c
			if v != N-1 {
				for _, P := range E[v] {
					c2 := c + P.x
					t2 := c2 % 28
					w := P.y
					if d[t2][w] == INF {
						heap.Push(pq, tuple{c2, t2, w})
					}
				}
			}
		}
	}

	ans := INF
	for i := 0; i < 28; i++ {
		if i%4 == 0 || i%7 == 0 {
			ans = min(ans, d[i][N-1])
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type tuple struct {
	c, t, v int
}

type Heap []tuple

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].c < h[j].c }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
