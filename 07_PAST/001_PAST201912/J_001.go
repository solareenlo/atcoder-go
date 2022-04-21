package main

import (
	"container/heap"
	"fmt"
)

const INF = 1 << 60

var (
	H, W int
	dx   = [4]int{0, 1, 0, -1}
	dy   = [4]int{-1, 0, 1, 0}
	A    = [50][50]int{}
)

func dijkstra(sx, sy int) [][]int {
	vis := make([]bool, H*W)
	res := make([][]int, H)
	for i := range res {
		res[i] = make([]int, W)
		for j := range res[i] {
			res[i][j] = INF
		}
	}
	que := &Heap{}

	res[sy][sx] = 0
	heap.Push(que, pair{0, sx + sy*W})

	for que.Len() > 0 {
		q := (*que)[0]
		heap.Pop(que)

		id := q.y

		if vis[id] {
			continue
		}
		vis[id] = true

		x := id % W
		y := id / W

		for d := 0; d < 4; d++ {
			xx := x + dx[d]
			yy := y + dy[d]
			if 0 <= xx && xx < W && 0 <= yy && yy < H {
				if res[yy][xx] > res[y][x]+A[yy][xx] {
					res[yy][xx] = res[y][x] + A[yy][xx]
					heap.Push(que, pair{res[yy][xx], xx + yy*W})
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Scan(&H, &W)
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			fmt.Scan(&A[y][x])
		}
	}

	mix := dijkstra(0, H-1)
	miy := dijkstra(W-1, H-1)
	miz := dijkstra(W-1, 0)

	ans := INF
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			ans = min(ans, mix[y][x]+miy[y][x]+miz[y][x]-2*A[y][x])
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type pair struct{ x, y int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x < h[j].x || (h[i].x == h[j].y && h[i].y < h[j].y) }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
