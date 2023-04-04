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
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	q := &HeapPair{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		heap.Push(q, pair{a, i})
	}

	ans := make([]pair, 0)
	for q.Len() > 1 {
		a := heap.Pop(q).(pair)
		val1 := a.x
		val2 := (*q)[0].x
		if clz(uint32(val1)) == clz(uint32(val2)) {
			val1 -= val2
			ans = append(ans, pair{a.y + 1, (*q)[0].y + 1})
			if val1 != 0 {
				heap.Push(q, pair{val1, a.y})
			}
			continue
		}
		if (val1 & 1) != 0 {
			ans = append(ans, pair{a.y + 1, 1})
			val1--
		}
		if val1 != 0 {
			ans = append(ans, pair{a.y + 1, a.y + 1})
			val1 >>= 1
			heap.Push(q, pair{val1, a.y})
		}
	}

	fmt.Fprintln(out, len(ans)+2)
	ans = reverseOrderPair(ans)
	fmt.Fprintln(out, 1, 1)
	for i := 0; i < len(ans); i++ {
		fmt.Fprintf(out, "2 %d %d\n", ans[i].x, ans[i].y)
	}
	fmt.Fprintln(out, 1, 1)
}

func reverseOrderPair(a []pair) []pair {
	n := len(a)
	res := make([]pair, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func clz(x uint32) int {
	return bits.LeadingZeros32(x)
}

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int { return len(h) }
func (h HeapPair) Less(i, j int) bool {
	if h[i].x == h[j].x {
		return h[i].y > h[j].y
	}
	return h[i].x > h[j].x
}
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
