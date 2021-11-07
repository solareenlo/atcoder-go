package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var (
	dx = [4]int{1, -1, 0, 0}
	dy = [4]int{0, 0, 1, -1}
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w, k, x3, y3, x2, y2 int
	fmt.Fscan(in, &h, &w, &k, &x3, &y3, &x2, &y2)
	x3--
	y3--
	x2--
	y2--

	s := make([]string, h)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}

	g := make([][]int, h)
	for i := range g {
		g[i] = make([]int, w)
		for j := range g[i] {
			g[i][j] = 1 << 60
		}
	}
	g[x3][y3] = 0

	q := &HEAP{}
	heap.Push(q, PAIR{0, x3, y3})
	for q.Len() > 0 {
		a := (*q)[0].a
		x := (*q)[0].x
		y := (*q)[0].y
		heap.Pop(q)
		if g[x][y] < a {
			continue
		}
		val := g[x][y] + 1
		for dir := 0; dir < 4; dir++ {
			xx := x
			yy := y
			for i := 1; i <= k; i++ {
				xx += dx[dir]
				yy += dy[dir]
				if xx < 0 || yy < 0 || xx >= h || yy >= w || s[xx][yy] == '@' || g[xx][yy] < val {
					break
				}
				if g[xx][yy] > val {
					g[xx][yy] = val
					heap.Push(q, PAIR{g[xx][yy], xx, yy})
				}
			}
		}
	}

	if g[x2][y2] == 1<<60 {
		fmt.Println(-1)
	} else {
		fmt.Println(g[x2][y2])
	}
}

type PAIR struct{ a, x, y int }
type HEAP []PAIR

func (h HEAP) Len() int            { return len(h) }
func (h HEAP) Less(i, j int) bool  { return h[i].a < h[j].a }
func (h HEAP) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HEAP) Push(x interface{}) { *h = append(*h, x.(PAIR)) }

func (h *HEAP) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
