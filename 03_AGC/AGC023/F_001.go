package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 200005

var (
	fa = make([]int, N)
)

func find(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	p := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}

	q := &Heap{}

	a := make([]Node, n+1)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		fa[i] = i
		tmp := 0
		if x == 0 {
			tmp = 1
		}
		a[i] = Node{tmp, x, i}
		heap.Push(q, a[i])
	}

	ans := 0
	vis := make([]int, N)
	vis[1] = 1
	for q.Len() > 0 {
		t := (*q)[0]
		x := t.id
		heap.Pop(q)
		if vis[x] == 0 && t.a == a[x].a && t.b == a[x].b {
			vis[x] = 1
			y := find(p[x])
			ans += a[y].b * t.a
			fa[x] = y
			a[y].a += t.a
			a[y].b += t.b
			heap.Push(q, a[y])
		}
	}
	fmt.Println(ans)
}

type Node struct{ a, b, id int }

type Heap []Node

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].a*h[j].b > h[i].b*h[j].a }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Node)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
