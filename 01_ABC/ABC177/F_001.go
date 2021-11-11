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

	var h, w int
	fmt.Fscan(in, &h, &w)

	bTree := bTreeNew(w)
	q := &Heap{}
	for i := 0; i < w; i++ {
		bTree.Add(i, 1)
		heap.Push(q, pair{0, i})
	}

	dp := make([]int, w)
	for i := 1; i < h+1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		if b != w && dp[b] == -1 {
			sum := bTree.Sum(b)
			if sum != 0 {
				k := bTree.Lowerbound(sum)
				dp[b] = b - k + dp[k]
				bTree.Add(b, 1)
				heap.Push(q, pair{dp[b], b})
			}
		}
		sum := bTree.Sum(a-2) + 1
		for {
			k := bTree.Lowerbound(sum)
			if k > b-1 {
				break
			}
			bTree.Add(k, -1)
			dp[k] = -1
		}
		for q.Len() != 0 && dp[(*q)[0].y] != (*q)[0].x {
			heap.Pop(q)
		}
		if q.Len() == 0 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, (*q)[0].x+i)
		}
	}
}

type binaryIndexedTree []int

func bTreeNew(n int) binaryIndexedTree {
	return make(binaryIndexedTree, n+1)
}

func (tree binaryIndexedTree) Sum(i int) int {
	sum := 0
	for i++; i > 0; i -= i & -i {
		sum += tree[i]
	}
	return sum
}
func (tree binaryIndexedTree) Add(i, x int) {
	for i++; i < len(tree) && i > 0; i += i & -i {
		tree[i] += x
	}
}

func (tree binaryIndexedTree) Lowerbound(x int) int {
	idx, k := 0, 1
	for k < len(tree) {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k < len(tree) && tree[idx+k] < x {
			x -= tree[idx+k]
			idx += k
		}
	}
	return idx
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
