package main

import (
	"container/list"
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([]string, h)
	d := make([][]int, h)
	for i := 0; i < h; i++ {
		d[i] = make([]int, w)
		fmt.Scan(&s[i])
		for j := 0; j < w; j++ {
			d[i][j] = 1 << 60
		}
	}
	d[0][0] = 0

	type pair struct{ x, y int }
	dq := list.New()
	dq.PushBack(pair{0, 0})
	for dq.Len() > 0 {
		x := dq.Front().Value.(pair).x
		y := dq.Front().Value.(pair).y
		dq.Remove(dq.Front())
		for dy := -2; dy <= 2; dy++ {
			for dx := -2; dx <= 2; dx++ {
				ny := y + dy
				nx := x + dx
				if !(0 <= ny && ny < h && 0 <= nx && nx < w) {
					continue
				}
				if abs(dx) == 2 && abs(dy) == 2 {
					continue
				}
				if s[ny][nx] == '.' && abs(dx)+abs(dy) == 1 && d[y][x] < d[ny][nx] {
					d[ny][nx] = d[y][x]
					dq.PushFront(pair{nx, ny})
				}
				if d[y][x]+1 < d[ny][nx] {
					d[ny][nx] = d[y][x] + 1
					dq.PushBack(pair{nx, ny})
				}
			}
		}
	}
	fmt.Println(d[h-1][w-1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
