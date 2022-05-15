package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

var pr = make([]int, 200005)

func f(p int) int {
	if p == pr[p] {
		return p
	}
	pr[p] = f(pr[p])
	return pr[p]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	gr := make([][]pair, 200005)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		gr[u] = append(gr[u], pair{v, w})
		gr[v] = append(gr[v], pair{u, w})
	}

	dt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dt[i] = 1 << 60
	}

	pq := &Heap{}
	for i := 1; i <= k; i++ {
		dt[i] = 0
		heap.Push(pq, pair{0, i})
		pr[i] = i
	}

	for pq.Len() > 0 {
		x := (*pq)[0]
		heap.Pop(pq)
		w := x.x
		u := x.y
		if w > dt[u] {
			continue
		}
		for _, y := range gr[u] {
			v := y.x
			c := y.y
			W := w + c
			if W < dt[v] {
				dt[v] = W
				pr[v] = pr[u]
				heap.Push(pq, pair{W, v})
			}
		}
	}

	type tuple struct{ x, y, z int }
	eg := make([]tuple, 200005)
	for i := 1; i <= n; i++ {
		for _, y := range gr[i] {
			j := y.x
			c := y.y
			if pr[i] < pr[j] {
				eg = append(eg, tuple{dt[i] + dt[j] + c, pr[i], pr[j]})
			}
		}
	}

	var q int
	fmt.Fscan(in, &q)
	u := make([]int, q)
	v := make([]int, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(in, &u[i], &v[i], &t)
		eg = append(eg, tuple{t, 1e9, i})
	}
	sort.Slice(eg, func(i, j int) bool {
		return eg[i].x < eg[j].x || (eg[i].x == eg[j].x && eg[i].y < eg[j].y) || (eg[i].x == eg[j].x && eg[i].y == eg[j].y && eg[i].z < eg[j].z)
	})

	ans := make([]bool, 200005)
	for _, z := range eg {
		x := z.y
		y := z.z
		if x < 1_000_000_000 {
			pr[f(x)] = f(y)
		} else if f(u[y]) == f(v[y]) {
			ans[y] = true
		}
	}

	for i := 0; i < q; i++ {
		if ans[i] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

type pair struct{ x, y int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
