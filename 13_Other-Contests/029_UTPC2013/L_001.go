package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 256

	var n int
	fmt.Fscan(in, &n)
	str := make([]string, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str[i])
	}

	a := make([][]int, N)
	for i := range a {
		a[i] = make([]int, N)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	dist := make([][][]int, N)
	for i := range dist {
		dist[i] = make([][]int, N)
		for j := range dist[i] {
			dist[i][j] = make([]int, 2)
			for k := range dist[i][j] {
				dist[i][j][k] = 1061109567
			}
		}
	}

	pq := &HeapNode{}
	var update func(int, int, int, int)
	update = func(s, e, x, m int) {
		if dist[s][e][m] > x {
			dist[s][e][m] = x
			heap.Push(pq, node{s, e, m, x})
		}
	}
	for i := 0; i < n; i++ {
		update(i, i, 0, 0)
	}
	for pq.Len() > 0 {
		x := heap.Pop(pq).(node)
		if x.m == 0 {
			for i := 0; i < n; i++ {
				update(x.s, i, x.cost+dist[x.e][i][0], 0)
				update(i, x.e, x.cost+dist[i][x.s][0], 0)
				if str[x.e][i] == '=' {
					update(x.s, i, x.cost+a[x.e][i], 0)
				}
				if str[i][x.s] == '=' {
					update(i, x.e, x.cost+a[i][x.s], 0)
				}
				if str[i][x.s] == '+' {
					update(i, x.e, x.cost+a[i][x.s], 1)
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if str[x.e][i] == '-' {
					update(x.s, i, x.cost+a[x.e][i], 0)
				}
			}
		}
	}
	if dist[0][n-1][0] > 1e9 {
		dist[0][n-1][0] = -1
	}
	fmt.Println(dist[0][n-1][0])
}

type node struct {
	s, e, m, cost int
}

type HeapNode []node

func (h HeapNode) Len() int            { return len(h) }
func (h HeapNode) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h HeapNode) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapNode) Push(x interface{}) { *h = append(*h, x.(node)) }

func (h *HeapNode) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
