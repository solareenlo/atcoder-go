package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type tuple struct {
		x, y, z int
	}

	var dx []int = []int{-1, 0, 0, 1}
	var dy []int = []int{0, -1, 1, 0}

	var H, W int
	fmt.Fscan(in, &H, &W)

	var d [505][505]int
	var a [505]string
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &a[i])
		for j := 0; j < W; j++ {
			d[i][j] = int(1e9)
		}
	}

	q := list.New()
	q.PushBack(tuple{0, 0, 0})
	d[0][0] = 0
	for q.Len() != 0 {
		tup := q.Front().Value.(tuple)
		dist := tup.x
		x := tup.y
		y := tup.z
		q.Remove(q.Front())
		if dist != d[x][y] {
			continue
		}
		for k := 0; k < 4; k++ {
			x2 := x + dx[k]
			y2 := y + dy[k]
			if x2 < 0 || x2 >= H || y2 < 0 || y2 >= W {
				continue
			}
			wt := false
			if a[x][y] != a[x2][y2] {
				wt = true
			}
			if !wt && d[x][y] < d[x2][y2] {
				d[x2][y2] = d[x][y]
				q.PushFront(tuple{d[x2][y2], x2, y2})
			} else if wt && d[x][y]+1 < d[x2][y2] {
				d[x2][y2] = d[x][y] + 1
				q.PushBack(tuple{d[x2][y2], x2, y2})
			}
		}
	}
	fmt.Println(d[H-1][W-1])
}
