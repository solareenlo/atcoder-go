package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N_MAX = 30
	const K_MAX = 10

	type pair struct {
		V, W int
	}

	var N, M, F, S int
	fmt.Fscan(in, &N, &M, &F, &S)
	F--
	S--
	g := make([][]int, N)
	rev := make([][]int, N)
	for M > 0 {
		M--
		var A, B int
		fmt.Fscan(in, &A, &B)
		A--
		B--
		g[A] = append(g[A], B)
		rev[B] = append(rev[B], A)
	}
	var K int
	fmt.Fscan(in, &K)
	jewels := make([]pair, K)
	jewel_idx := make([]int, N)
	for i := range jewel_idx {
		jewel_idx[i] = -1
	}
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &jewels[i].V, &jewels[i].W)
		jewels[i].V--
		jewel_idx[jewels[i].V] = i
	}

	var dp [1 << K_MAX][N_MAX][N_MAX]int
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -2139062144
			}
		}
	}
	for i := range dp[0] {
		for j := range dp[0][i] {
			dp[0][i][j] = 0x00
		}
	}
	for bit := 1; bit < 1<<K; bit++ {
		jewel := make([]int, N)
		for i := 0; i < K; i++ {
			if (bit & (1 << i)) != 0 {
				jewel[jewels[i].V] = jewels[i].W
			}
		}
		var deg [N_MAX][N_MAX]int
		for x := 0; x < N; x++ {
			for y := 0; y < N; y++ {
				deg[x][y] = len(g[x])
			}
		}

		q := &HeapTuple{}

		var decide func(int, int)
		var update func(int, int, int)

		decide = func(x, y int) {
			c := -dp[bit][x][y]
			for _, w := range rev[y] {
				if jewel[w] == 0 {
					update(w, x, c)
				}
			}
		}

		update = func(x, y, c int) {
			deg[x][y]--
			if deg[x][y] == 0 {
				if dp[bit][x][y] < c {
					dp[bit][x][y] = c
				}
				decide(x, y)
			}
			if dp[bit][x][y] >= c {
				return
			}
			dp[bit][x][y] = c
			heap.Push(q, Tuple{c, x, y})
		}

		for x := 0; x < N; x++ {
			if jewel[x] == 0 {
				for y := 0; y < N; y++ {
					if jewel[y] == 0 {
						for _, z := range g[x] {
							if jewel[z] != 0 {
								bit2 := bit ^ (1 << jewel_idx[z])
								update(x, y, jewel[z]-dp[bit2][y][z])
							}
						}
					}
				}
			}
		}

		for q.Len() > 0 {
			tmp := heap.Pop(q).(Tuple)
			c := tmp.x
			x := tmp.y
			y := tmp.z
			if c <= 0 {
				break
			}
			if deg[x][y] <= 0 {
				continue
			}
			deg[x][y] = 0
			decide(x, y)
		}

		for x := 0; x < N; x++ {
			if jewel[x] == 0 {
				for y := 0; y < N; y++ {
					if jewel[y] == 0 {
						if dp[bit][x][y] == -2139062144 {
							dp[bit][x][y] = 0
						}
					}
				}
			}
		}
	}

	fmt.Println(dp[(1<<K)-1][F][S])
}

type Tuple struct {
	x, y, z int
}

type HeapTuple []Tuple

func (h HeapTuple) Len() int            { return len(h) }
func (h HeapTuple) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapTuple) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTuple) Push(x interface{}) { *h = append(*h, x.(Tuple)) }

func (h *HeapTuple) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
