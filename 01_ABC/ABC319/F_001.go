package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 510
	var t, g, l, s [N]int
	var to [N][]int
	var u [10]int
	var v [10]bool
	var f [1024]int

	var n int
	fmt.Fscan(in, &n)
	mx := 0
	m := 0
	for i := 2; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x, &t[i], &s[i], &g[i])
		to[x] = append(to[x], i)
		if t[i] == 2 {
			l[i] = m
			u[m] = i
			m++
		} else {
			mx = max(mx, s[i])
		}
	}
	f[0], t[1] = 1, 1
	q := &HeapPair{}
	for i := 0; i < 1<<m; i++ {
		if f[i] == 0 {
			continue
		}
		for q.Len() > 0 {
			heap.Pop(q)
		}
		q.Push(pair{1, 1})
		ct, S := 0, 0
		for j := 0; j < m; j++ {
			v[j] = false
		}
		for q.Len() > 0 {
			tmp := heap.Pop(q).(pair)
			x := -tmp.x
			y := tmp.y
			ct++
			if x > f[i] {
				break
			}
			if t[y] == 1 {
				f[i] += g[y]
				S += g[y]
			} else if ((i>>l[y])&1)^1 != 0 {
				v[l[y]] = true
				continue
			}
			for _, z := range to[y] {
				if t[z] == 2 {
					heap.Push(q, pair{1, z})
				} else {
					heap.Push(q, pair{-s[z], z})
				}
			}
		}
		if ct == n || f[i] >= mx {
			fmt.Println("Yes")
			return
		}
		for j := 0; j < m; j++ {
			if v[j] {
				f[i|(1<<j)] = max(f[i|(1<<j)], f[i]*g[u[j]]-S)
			}
		}
	}
	fmt.Println("No")
}

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
