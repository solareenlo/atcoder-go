package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const MAX_H = 500
const MAX_W = 500
const MAX_X = MAX_H*MAX_W - 2
const INF = int(1e18)

var dx [4]int = [4]int{1, 0, -1, 0}
var dy [4]int = [4]int{0, 1, 0, -1}

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W, X, sx, sy, gx, gy int
	fmt.Fscan(in, &H, &W, &X, &sx, &sy, &gx, &gy)
	sx--
	sy--
	gx--
	gy--
	var A [MAX_H][MAX_W]int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	var C [MAX_X + 1]int
	for i := 0; i < X; i++ {
		fmt.Fscan(in, &C[i+1])
	}
	var dst [MAX_H][MAX_W]int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			dst[i][j] = INF
		}
	}
	pq := &HeapTuple{}
	dst[sx][sy] = 0
	heap.Push(pq, tuple{sx, sy, 0})
	for pq.Len() > 0 {
		s := heap.Pop(pq).(tuple)
		if dst[s.x][s.y] < s.dst {
			continue
		}
		for i := 0; i < 4; i++ {
			x2 := s.x + dx[i]
			y2 := s.y + dy[i]
			if x2 < 0 || x2 >= H || y2 < 0 || y2 >= W {
				continue
			}
			cst := 0
			if A[x2][y2] != A[s.x][s.y] {
				cst = C[A[x2][y2]]
			}
			if dst[x2][y2] > s.dst+cst {
				dst[x2][y2] = s.dst + cst
				heap.Push(pq, tuple{x2, y2, dst[x2][y2]})
			}
		}
	}
	fmt.Println(dst[gx][gy])
}

type tuple struct {
	x, y, dst int
}

type HeapTuple []tuple

func (h HeapTuple) Len() int            { return len(h) }
func (h HeapTuple) Less(i, j int) bool  { return h[i].dst < h[j].dst }
func (h HeapTuple) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTuple) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *HeapTuple) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
