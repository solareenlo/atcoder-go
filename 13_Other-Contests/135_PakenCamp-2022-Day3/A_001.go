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
	N := 2*n + 1
	A := make([][]int, N)
	for i := range A {
		A[i] = make([]int, N)
	}
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
		for j := range dp[i] {
			dp[i][j] = 1 << 61
		}
	}
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			fmt.Fscan(in, &A[y][x])
			if A[y][x] == -1 {
				A[y][x] = 1 << 60
			}
		}
	}
	pq := &HeapTuple{}
	for y := 0; y < N; y++ {
		for x := 0; x < n; x++ {
			if y == 0 || x == 0 || y == N-1 {
				dp[y][x] = A[y][x]
				heap.Push(pq, tuple{A[y][x], y, x})
			}
		}
	}
	type pair struct {
		x, y int
	}
	dir := []pair{
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {0, 1}, {1, 0}, {-1, 0}, {0, -1},
	}
	for pq.Len() != 0 {
		tmp := heap.Pop(pq).(tuple)
		d := tmp.x
		y := tmp.y
		x := tmp.z
		if d > dp[y][x] {
			continue
		}
		for i := 0; i < 8; i++ {
			dy := dir[i].x
			dx := dir[i].y
			ny := y + dy
			nx := x + dx
			if ny < 0 || nx < 0 || ny >= N || nx >= N {
				continue
			}
			if d+A[ny][nx] >= dp[ny][nx] {
				continue
			}
			dp[ny][nx] = d + A[ny][nx]
			heap.Push(pq, tuple{dp[ny][nx], ny, nx})
		}
	}
	ans := 1 << 61
	for y := 0; y < N; y++ {
		for x := n + 1; x < N; x++ {
			if y == 0 || x == N-1 || y == N-1 {
				ans = min(ans, dp[y][x])
			}
		}
	}
	fmt.Println(ans)
}

type tuple struct {
	x, y, z int
}

type HeapTuple []tuple

func (h HeapTuple) Len() int            { return len(h) }
func (h HeapTuple) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapTuple) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTuple) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *HeapTuple) Pop() interface{} {
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
