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

	const INF = 4557430888798830399

	dir := "DURL"

	var sum [4][400000]int
	var s [2000]string
	var dx [4]int = [4]int{1, -1, 0, 0}
	var dy [4]int = [4]int{0, 0, 1, -1}

	var n, m int
	fmt.Fscan(in, &n, &m)
	var k int
	fmt.Fscan(in, &k)
	var t string
	fmt.Fscan(in, &t)
	t += t
	for i := 0; i < 4; i++ {
		for j := 0; j < len(t); j++ {
			sum[i][j+1] += sum[i][j]
			if t[j] == dir[i] {
				sum[i][j+1]++
			}
		}
	}
	sx, sy, gx, gy := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		for j := 0; j < m; j++ {
			if s[i][j] == 'S' {
				sx = i
				sy = j
			}
			if s[i][j] == 'G' {
				gx = i
				gy = j
			}
		}
	}
	var d [2000][2000]int
	for i := range d {
		for j := range d[i] {
			d[i][j] = INF
		}
	}
	que := make(HeapTuple, 0)
	d[sx][sy] = 0
	heap.Init(&que)
	heap.Push(&que, tuple{sx, sy, 0})
	for que.Len() > 0 {
		p := heap.Pop(&que).(tuple)
		if d[p.a][p.b] != p.c {
			continue
		}
		l := p.c % k
		for i := 0; i < 4; i++ {
			nx := p.a + dx[i]
			ny := p.b + dy[i]
			if nx < 0 || nx >= n || ny < 0 || ny >= m || s[nx][ny] == '#' {
				continue
			}
			a := upperBound(sum[i][:len(t)+1], sum[i][l]) - 1
			if a == len(t) {
				continue
			}
			if d[nx][ny] > p.c+(a-l)+1 {
				d[nx][ny] = p.c + (a - l) + 1
				heap.Push(&que, tuple{nx, ny, p.c + (a - l) + 1})
			}
		}
	}
	if d[gx][gy] == INF {
		fmt.Println(-1)
		return
	}
	fmt.Println(d[gx][gy])
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

type tuple struct {
	a, b, c int
}

type HeapTuple []tuple

func (h HeapTuple) Len() int            { return len(h) }
func (h HeapTuple) Less(i, j int) bool  { return h[i].c < h[j].c }
func (h HeapTuple) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTuple) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *HeapTuple) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
