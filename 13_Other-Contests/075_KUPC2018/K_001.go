package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

const Q = 100000

var n, m int
var a, b, l, r, ord []int

func main() {
	in := bufio.NewReader(os.Stdin)

	var tmp int
	fmt.Fscan(in, &n, &m, &tmp)
	a = make([]int, n)
	b = make([]int, m)
	l = make([]int, n)
	r = make([]int, n)
	ord = make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		l[i]--
		r[i]--
	}
	for i := range ord {
		ord[i] = i
	}
	sort.Slice(ord, func(i, j int) bool {
		return l[ord[i]] < l[ord[j]]
	})
	ok := 0
	ng := Q*Q + 1
	for ng-ok > 1 {
		mid := (ok + ng) >> 1
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	fmt.Println(float64(ok) / float64(Q))
}

func check(k int) bool {
	color := 0
	color++
	usedA := make([]int, n)
	usedB := make([]int, m)
	weiA := make([]int, n)
	weiB := make([]int, m)
	val := 0
	pq := &HeapPair{}
	for i := 0; i < n; i++ {
		weiA[i] = k - Q*a[i]
		if weiA[i] <= 0 {
			usedA[i] = color
			val += weiA[i]
		}
	}
	for i := 0; i < m; i++ {
		weiB[i] = k - Q*b[i]
		if weiB[i] <= 0 {
			usedB[i] = color
			val += weiB[i]
		}
	}
	now := 0
	head := 0
	for {
		for now < m && usedB[now] == color {
			now++
		}
		if now == m {
			break
		}
		for head < n && l[ord[head]] <= now {
			if usedA[ord[head]] != color {
				heap.Push(pq, pair{r[ord[head]], ord[head]})
			}
			head++
		}
		if pq.Len() == 0 {
			now++
			continue
		}

		tmp := heap.Pop(pq).(pair)
		rv := tmp.x
		idx := tmp.y
		if rv < now {
			continue
		}
		use := min(weiB[now], weiA[idx])
		weiB[now] -= use
		weiA[idx] -= use
		val += use
		if weiA[idx] > 0 && now+1 <= rv {
			heap.Push(pq, pair{rv, idx})
		}
		if weiB[now] == 0 {
			now++
		}
	}
	return val <= 0
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
