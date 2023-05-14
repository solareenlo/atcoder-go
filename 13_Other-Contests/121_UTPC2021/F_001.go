package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k1, k2 int
	fmt.Fscan(in, &n, &k1, &k2)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h[i])
	}

	N := 200010
	var s1, s2, sr1, sr2 FastSet
	s1.init(N)
	s2.init(N)
	sr1.init(N)
	sr2.init(N)
	s1.insert(0)
	s2.insert(0)
	for i := 0; i < N; i++ {
		sr1.insert(i)
		sr2.insert(i)
	}

	a1 := make([]int, n+1)
	a2 := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		a1[i] = -1
		a2[i] = -1
	}
	var f func([]int, int)
	f = func(a []int, k int) {
		pq := &Heap{}
		for i := n - 1; i >= 0; i-- {
			heap.Push(pq, h[i])
			if pq.Len() > k {
				heap.Pop(pq)
			}
			if pq.Len() == k {
				a[i] = heap.Pop(pq).(int)
			}
		}
	}
	f(a1, k2)
	f(a2, k1)

	dp0 := 0
	for i := 0; i < n; i++ {
		for sr1.min() <= a1[i] || sr2.min() <= a1[i] {
			var x int
			if sr2.min() <= a1[i] {
				x = sr2.min()
				sr2.erase(x)
				s2.insert(x)
			} else {
				x = sr1.min()
				sr1.erase(x)
				s1.insert(x)
			}
			a := s1.next(a1[i] + 1)
			b := s2.next(a1[i] + 1)
			if a == b && a < N {
				s2.erase(a)
				sr2.insert(a)
			} else if a < b {
				s1.erase(a)
				sr1.insert(a)
			} else if a > b {
				s2.erase(b)
				sr2.insert(b)
			}
		}
		for sr1.min() <= a2[i] {
			x := sr1.min()
			sr1.erase(x)
			s1.insert(x)
			a := s1.next(a2[i] + 1)
			b := s2.next(x + 1)
			if a == b && a < N {
				s2.erase(a)
				sr2.insert(a)
			} else if a < b {
				s1.erase(a)
				sr1.insert(a)
			} else if a > b {
				s2.erase(b)
				sr2.insert(b)
			}
		}
		for _, e := range []int{h[i] - k1, h[i] - k2} {
			if e < 0 {
				continue
			}
			dp0++
			x := s1.next(e + 1)
			y := s2.next(e + 1)
			if x == y && x < N {
				s2.erase(x)
				sr2.insert(x)
			} else if x < y {
				s1.erase(x)
				sr1.insert(x)
			} else if x > y {
				s2.erase(y)
				sr2.insert(y)
			}
		}
	}
	for s1.min() < N {
		s1.erase(s1.min())
		dp0++
	}
	for s2.min() < N {
		s2.erase(s2.min())
		dp0++
	}
	fmt.Println(dp0 - 2)
}

// 0 以上 2 ^ (6D) 未満の整数を取れる
// D : 3 -> 2.6 * 10 ^ 5
// D : 4 -> 1.6 * 10 ^ 7
// D : 5 -> 1.0 * 10 ^ 9
// D : 6 -> int の範囲全部
const D = 3

type FastSet struct {
	n int
	a [D][]uint64
}

func (set *FastSet) init(n_ int) {
	set.n = n_
	for i := 0; i < D; i++ {
		n_ = (n_ + 63) >> 6
		set.a[i] = make([]uint64, n_)
	}
}

func (set FastSet) empty() bool { return set.a[D-1][0] == 0 }

func (set FastSet) contains(x int) bool {
	return ((set.a[0][x>>6] >> (x & 63)) & 1) == 1
}

func (set *FastSet) insert(x int) {
	for i := 0; i < D; i++ {
		y := x & 63
		x >>= 6
		set.a[i][x] |= uint64(1 << y)
	}
}

func (set *FastSet) erase(x int) {
	for i := 0; i < D; i++ {
		y := x & 63
		x >>= 6
		set.a[i][x] &= ^uint64(1 << y)
		if set.a[i][x] != 0 {
			break
		}
	}
}

func (set FastSet) next(x int) int {
	for i := 0; i < D; i++ {
		k := x >> 6
		y := x & 63
		if k >= len(set.a[i]) {
			return set.n
		}
		top := set.a[i][k] >> y
		if top != 0 {
			x += ctz(uint64(top))
			for j := i - 1; j >= 0; j-- {
				x = x<<6 | ctz(set.a[j][x])
			}
			return x
		}
		x = k + 1
	}
	return set.n
}

func (set FastSet) prev(x int) int {
	for i := 0; i < D; i++ {
		if x < 0 {
			return -1
		}
		k := x >> 6
		y := x & 63
		bot := set.a[i][k] << (63 - y)
		if bot != 0 {
			x -= clz(bot)
			for j := i - 1; j >= 0; j-- {
				x = x<<6 | (63 - clz(set.a[j][x]))
			}
			return x
		}
		x = k - 1
	}
	return -1
}

func clz(x uint64) int {
	return bits.LeadingZeros64(x)
}

func ctz(x uint64) int {
	return bits.TrailingZeros64(x)
}

func (set FastSet) max() int { return set.prev(set.n) }
func (set FastSet) min() int { return set.next(0) }

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
