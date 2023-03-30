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

	var H, W, n, f, sx, sy, gx, gy int
	fmt.Fscan(in, &H, &W, &n, &f, &sx, &sy, &gx, &gy)
	tate := make([][]pair, H+1)
	yoko := make([][]pair, W+1)
	Go := make([][]pair, n*7+2)
	sta := -1
	for i := 0; i < n; i++ {
		var a, b, d, e int
		var c string
		fmt.Fscan(in, &a, &b, &c, &d, &e)
		if sx == a && sy == b {
			sta = i
		}
		tate[a] = append(tate[a], pair{b, i + n})
		Go[i+n] = append(Go[i+n], pair{0, i})
		yoko[b] = append(yoko[b], pair{a, i + n*2})
		Go[i+n*2] = append(Go[i+n*2], pair{0, i})
		tate[a] = append(tate[a], pair{b - d, i + n*3})
		if c == "W" {
			Go[i] = append(Go[i], pair{0, i + n*3})
		} else {
			Go[i] = append(Go[i], pair{e, i + n*3})
		}
		tate[a] = append(tate[a], pair{b + d, i + n*4})
		if c == "E" {
			Go[i] = append(Go[i], pair{0, i + n*4})
		} else {
			Go[i] = append(Go[i], pair{e, i + n*4})
		}
		yoko[b] = append(yoko[b], pair{a - d, i + n*5})
		if c == "N" {
			Go[i] = append(Go[i], pair{0, i + n*5})
		} else {
			Go[i] = append(Go[i], pair{e, i + n*5})
		}
		yoko[b] = append(yoko[b], pair{a + d, i + n*6})
		if c == "S" {
			Go[i] = append(Go[i], pair{0, i + n*6})
		} else {
			Go[i] = append(Go[i], pair{e, i + n*6})
		}
	}
	if sta == -1 {
		fmt.Println(-1)
		return
	}
	tate[gx] = append(tate[gx], pair{gy, n * 7})
	yoko[gy] = append(yoko[gy], pair{gx, n*7 + 1})
	for i := 1; i <= H; i++ {
		if len(tate[i]) == 0 {
			continue
		}
		sort.Slice(tate[i], func(a, b int) bool {
			if tate[i][a].x == tate[i][b].x {
				return tate[i][a].y < tate[i][b].y
			}
			return tate[i][a].x < tate[i][b].x
		})
		mae := tate[i][0].y
		for j := 1; j < len(tate[i]); j++ {
			it := tate[i][j].y
			cost := (tate[i][j].x - tate[i][j-1].x) * f
			Go[mae] = append(Go[mae], pair{cost, it})
			Go[it] = append(Go[it], pair{cost, mae})
			mae = it
		}
	}
	for i := 1; i <= W; i++ {
		if len(yoko[i]) == 0 {
			continue
		}
		sort.Slice(yoko[i], func(a, b int) bool {
			if yoko[i][a].x == yoko[i][b].x {
				return yoko[i][a].y < yoko[i][b].y
			}
			return yoko[i][a].x < yoko[i][b].x
		})
		mae := yoko[i][0].y
		for j := 1; j < len(yoko[i]); j++ {
			it := yoko[i][j].y
			cost := (yoko[i][j].x - yoko[i][j-1].x) * f
			Go[mae] = append(Go[mae], pair{cost, it})
			Go[it] = append(Go[it], pair{cost, mae})
			mae = it
		}
	}
	// 単純ダイクストラ法
	que := &HeapPair{}
	heap.Push(que, pair{0, sta})
	dis := make([]int, 7*n+2)
	for i := range dis {
		dis[i] = int(1e18)
	}
	dis[sta] = 0
	for que.Len() > 0 {
		tmp := heap.Pop(que).(pair)
		time := tmp.x
		ter := tmp.y
		if ter >= 7*n {
			fmt.Println(time)
			return
		}
		if dis[ter] < time {
			continue
		}
		for _, it := range Go[ter] {
			if dis[it.y] > time+it.x {
				dis[it.y] = time + it.x
				heap.Push(que, pair{time + it.x, it.y})
			}
		}
	}
	fmt.Println(-1)
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
