package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w, sx, sy int
	fmt.Scan(&h, &w, &sx, &sy)
	sx--
	sy--
	B := make([][]int, h)
	for i := range B {
		B[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Scan(&B[i][j])
		}
	}

	ans := make([][]string, h)
	for i := range ans {
		ans[i] = make([]string, w)
		for j := range ans[i] {
			ans[i][j] = "."
		}
	}
	dx := []int{0, -1, 0, 1}
	dy := []int{1, 0, -1, 0}
	x, y := sx, sy
	for {
		ans[x][y] = "W"
		next := -1
		for k := 0; k < 4; k++ {
			x2 := x + dx[k]
			y2 := y + dy[k]
			if 0 <= x2 && x2 < h && 0 <= y2 && y2 < w && B[x2][y2] < B[x][y] {
				if next == -1 || B[x2][y2] < B[x+dx[next]][y+dy[next]] {
					next = k
				}
			}
		}
		if next == -1 {
			break
		}
		x += dx[next]
		y += dy[next]
	}
	for i := 0; i < h; i++ {
		fmt.Println(strings.Join(ans[i], ""))
	}
}
