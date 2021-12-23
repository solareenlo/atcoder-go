package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type edge struct{ v, w, nxt int }

const NN = 200010

var (
	n, m int
	aa   = [NN]int{}
	b    = [NN]int{}
	dis  = [NN << 4]int{}
	visi = [NN << 4]bool{}
	e    = [NN << 4]edge{}
	tot  int
	hh   = [NN << 4]int{}
	v    = make([]int, 0)
	vv   = make([]int, 0)
)

func add(u, v, w int) {
	tot++
	e[tot] = edge{v, w, hh[u]}
	hh[u] = tot
}

func dijkstra() {
	q := &Heap{}
	heap.Push(q, pair{0, 1})
	dis[1] = 0
	for i := range dis {
		dis[i] = 1 << 60
	}
	for q.Len() > 0 {
		w := -(*q)[0].x
		u := (*q)[0].y
		heap.Pop(q)
		if visi[u] {
			continue
		}
		visi[u] = true
		for i := hh[u]; i > 0; i = e[i].nxt {
			v := e[i].v
			if w+e[i].w < dis[v] {
				dis[v] = w + e[i].w
				heap.Push(q, pair{-dis[v], v})
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &aa[i])
		aa[i] %= m
		v = append(v, m-aa[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		b[i] %= m
		v = append(v, b[i])
	}
	v = unique(v)
	for i := 1; i <= n; i++ {
		add(i, lowerBound(v, m-aa[i])+n+1, 0)
		add(lowerBound(v, b[i])+n+1, i, 0)
	}
	for i := 0; i < len(v)-1; i++ {
		add(i+n+1, i+n+2, v[i+1]-v[i])
	}
	add(len(v)+n, n+1, (v[0]-v[len(v)-1]+m)%m)
	dijkstra()
	fmt.Println(dis[n])
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
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
	sort.Ints(result)
	return result
}

type pair struct{ x, y int }
type Heap []pair

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i].x > h[j].x
}
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
