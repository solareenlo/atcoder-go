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

	var n, m, k, l int
	fmt.Fscan(in, &n, &m, &k, &l)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	pq := &Heap{}
	for i := 0; i < l; i++ {
		var u int
		fmt.Fscan(in, &u)
		heap.Push(pq, tuple{0, u - 1, a[u-1]})
	}

	type pair struct{ x, y int }
	e := make([][]pair, n)
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		e[u-1] = append(e[u-1], pair{v - 1, c})
		e[v-1] = append(e[v-1], pair{u - 1, c})
	}

	cnt := make([]int, n)
	d := make([]int, n)
	d2 := make([]int, n)
	for i := range d {
		d[i] = -1
		d2[i] = -1
	}
	country := make([]int, n)
	for pq.Len() > 0 {
		z := (*pq)[0].x
		x := (*pq)[0].y
		y := (*pq)[0].z
		heap.Pop(pq)
		if cnt[x] == 0 {
			d[x] = -z
			country[x] = y
		} else if (cnt[x] == 1) && (y != country[x]) {
			d2[x] = -z
		} else {
			continue
		}
		cnt[x]++
		sz := len(e[x])
		for i := 0; i < sz; i++ {
			heap.Push(pq, tuple{z - e[x][i].y, e[x][i].x, y})
		}
	}

	for i := 0; i < n; i++ {
		if a[i] != country[i] {
			fmt.Fprint(out, d[i])
		} else {
			fmt.Fprint(out, d2[i])
		}
		if i < (n - 1) {
			fmt.Fprint(out, " ")
		} else {
			fmt.Fprintln(out)
		}
	}
}

type tuple struct{ x, y, z int }
type Heap []tuple

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
