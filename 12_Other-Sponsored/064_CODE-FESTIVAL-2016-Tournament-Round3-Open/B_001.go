package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] -= 1
	}
	var s string
	fmt.Fscan(in, &s)

	lw, up := 0, n
	for up-lw > 1 {
		mid := (lw + up) / 2
		val := make([]int, n)
		for i := 0; i < n; i++ {
			if a[i] >= mid {
				val[i] = 1
			}
		}
		cnt := make([]int, n)
		u := make([]int, n)
		nx := make([]int, n+1)
		pr := make([]int, n+1)
		for i := 0; i < n+1; i++ {
			nx[i] = -1
			pr[i] = -1
		}
		root := n
		nx[root] = root
		pr[root] = root
		var link func(int, int)
		link = func(i, j int) {
			nx[i] = j
			pr[j] = i
		}
		var insert func(int, int)
		insert = func(i, j int) {
			link(j, nx[i])
			link(i, j)
		}
		var erase func(int)
		erase = func(i int) {
			link(pr[i], nx[i])
		}
		pre := -1
		for i := 0; i < n; i++ {
			if val[i] != pre {
				insert(pr[root], i)
				u[i] = 1
				pre = val[i]
			}
			cnt[pr[root]]++
		}
		var off [2]int
		pq := make([]HeapPair, 2)
		for i := 0; i < n; i++ {
			if u[i] != 0 {
				heap.Push(&pq[val[i]], pair{cnt[i], i})
			}
		}

		var modify_edge func(int, int)
		modify_edge = func(i, k int) {
			if val[i] == k {
				cnt[i]--
				heap.Push(&pq[val[i]], pair{cnt[i], i})
			}
		}

		rem := 0
		for i := range u {
			rem += u[i]
		}
		var del func(int)
		del = func(i int) {
			u[i] = 0
			rem--
			erase(i)
		}

		for _, op := range s {
			k := 0
			if op == 'M' {
				k = 1
			}
			if rem == 1 {
				break
			}
			modify_edge(nx[root], k)
			modify_edge(pr[root], k)
			off[k]++
			off[k^1]--
			for pq[k^1].Len() > 0 {
				tmp := pq[k^1][0]
				c := tmp.x
				i := tmp.y
				if cnt[i] != c || u[i] == 0 {
					heap.Pop(&pq[k^1])
					continue
				}
				c += off[k^1]
				if c == 0 {
					p, q := pr[i], nx[i]
					if p != root && q != root {
						cnt[p] += cnt[q] + off[val[q]]
						heap.Push(&pq[val[p]], pair{cnt[p], p})
						del(q)
					}
					del(i)
					heap.Pop(&pq[k^1])
				} else {
					break
				}
			}
		}
		if val[nx[root]] != 0 {
			lw = mid
		} else {
			up = mid
		}
	}
	fmt.Println(lw + 1)
}

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
