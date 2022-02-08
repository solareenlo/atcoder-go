package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var (
	a  = make([]int, 200200)
	m  = make([][20]int, 200200)
	pq = &Heap{}
)

func mn(p, q int) int {
	if a[p] < a[q] {
		return p
	}
	return q
}

func f(l, r int) int {
	j := 0
	for 1<<(j+1) <= r-l+1 {
		j++
	}
	return mn(m[l][j], m[r-(1<<j)+1][j])
}

func add(l, r int) {
	var x, y int
	if l > r {
		return
	}
	x = f(l, r)
	if x+1 == r {
		y = r
	} else {
		y = mn(f(x+1, r-1), r)
	}
	heap.Push(pq, st{l, r, x, y})
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		m[i][1] = i
	}

	for j := 2; ; j++ {
		k := 1 << j
		if k > n {
			break
		}
		for i := 0; i+k <= n; i++ {
			m[i][j] = mn(m[i][j-1], m[i+k/2][j-1])
		}
	}

	add(0, n-1)
	for pq.Len() > 0 {
		t := (*pq)[0]
		heap.Pop(pq)
		fmt.Fprint(out, a[t.x], " ", a[t.y], " ")
		add(t.l, t.x-1)
		add(t.x+1, t.y-1)
		add(t.y+1, t.r)
	}
}

type st struct{ l, r, x, y int }
type Heap []st

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return a[h[i].x] < a[h[j].x] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(st)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
