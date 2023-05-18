package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q, k int
	fmt.Fscan(in, &n, &q, &k)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	res := sol(n, k, v)
	for i := 0; i < q; i++ {
		var s, t int
		fmt.Fscan(in, &s, &t)
		s--
		t--
		if res[s].x <= t && t <= res[s].y {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

type Pair struct {
	x, y int
}

func sol(n, k int, v []int) []Pair {
	res := make([]Pair, n)
	tmp := sol2(n, k, v)
	v = ReverseOrderInt(v)
	tmp2 := sol2(n, k, v)
	for i := 0; i < n; i++ {
		res[i] = Pair{tmp[i], n - 1 - tmp2[n-1-i]}
	}
	return res
}

func sol2(n, k int, v []int) []int {
	res := make([]int, n)
	st := make([]Pair, 0)
	pq := &HeapPair{}
	var bit BIT
	bit.init(n)
	for i := 0; i < n; i++ {
		bit.add(i, 1)
	}
	for i := 0; i < n; i++ {
		l := lowerBoundPair(st, Pair{-(k - v[i]), -1}) - 1
		for pq.Len() > 0 && (*pq)[0].x+v[i] > k {
			bit.add(heap.Pop(pq).(Pair).y, -1)
		}
		heap.Push(pq, Pair{v[i], i})
		if l < 0 {
			res[i] = 0
		} else {
			l = st[l].y
			res[i] = l + 1 - bit.get(0, l+1)
		}
		for len(st) != 0 && st[len(st)-1].x >= -v[i] {
			st = st[:len(st)-1]
		}
		st = append(st, Pair{-v[i], i})
	}
	return res
}

type BIT struct {
	bit []int
	n   int
}

func (b *BIT) init(n int) {
	b.n = n
	b.bit = make([]int, n+1)
}

func (b BIT) sum(i int) int {
	res := 0
	for ; i > 0; i -= i & -i {
		res += b.bit[i]
	}
	return res
}

func (b *BIT) add(i, x int) {
	for i = i + 1; i <= b.n; i += i & -i {
		b.bit[i] += x
	}
}

func (bit BIT) get(a, b int) int {
	if b <= a {
		return 0
	}
	return bit.sum(b) - bit.sum(a)
}

func (b BIT) solve(x int) int { // sum([0,r])>=x となる最小のr
	if x <= 0 {
		return -1
	}
	l := 0
	Len := 1
	for Len < b.n {
		Len <<= 1
	}
	for ; Len > 0; Len >>= 1 {
		if l+Len < b.n && b.bit[l+Len] < x {
			x -= b.bit[l+Len]
			l += Len
		}
	}
	return l
}

type HeapPair []Pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(Pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func ReverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func lowerBoundPair(a []Pair, x Pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		if a[i].x == x.x {
			return a[i].y >= x.y
		}
		return a[i].x >= x.x
	})
	return idx
}
