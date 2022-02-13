package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const mod = 1_000_000_007

var (
	N int
	G = make([][]pair, 1<<17)
	a = make([]int, 1<<17)
	b = make([]int, 1<<17)
	c = make([]int, 1<<17)
	d = make([]int, 1<<17)
	P = &Heap{}
)

func calc(S int, d, w []int) {
	w[S] = 1
	for i := 1; i <= N; i++ {
		d[i] = -1e18
	}
	d[S] = 0
	heap.Push(P, pair{d[S], S})
	for P.Len() > 0 {
		c := (*P)[0].x
		u := (*P)[0].y
		heap.Pop(P)
		if d[u] == c {
			for _, p := range G[u] {
				v := p.x
				nxt := c - p.y
				if d[v] < nxt {
					w[v] = w[u]
					d[v] = nxt
					heap.Push(P, pair{d[v], v})
				} else if d[v] == nxt {
					w[v] += w[u]
					w[v] %= mod
				}
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var M, S, T int
	fmt.Fscan(in, &N, &M, &S, &T)
	for i := 0; i < M; i++ {
		var u, v, D int
		fmt.Fscan(in, &u, &v, &D)
		G[u] = append(G[u], pair{v, D})
		G[v] = append(G[v], pair{u, D})
	}

	calc(S, a, b)
	calc(T, c, d)

	ans := b[T] * b[T]
	t := a[T]
	for u := 0; u < N; {
		u++
		if 2*a[u] == t {
			ans -= b[u] * b[u] % mod * d[u] % mod * d[u] % mod
		}
		for _, p := range G[u] {
			v := p.x
			if 2*a[u] > t && 2*a[v] < t && a[u]+c[v]-p.y == t {
				ans -= b[u] * b[u] % mod * d[v] % mod * d[v] % mod
			}
		}
	}
	fmt.Println(ans % mod)
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
