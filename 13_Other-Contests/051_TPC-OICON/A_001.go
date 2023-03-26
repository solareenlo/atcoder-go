package main

import (
	"fmt"
)

const INF = 1 << 29

var dx = []int{0, -1, 0, 1}
var dy = []int{1, 0, -1, 0}

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	B := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&B[i])
	}

	var sx, sy, gx, gy int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if B[i][j] == 'S' {
				sx, sy = i, j
			}
			if B[i][j] == 'G' {
				gx, gy = i, j
			}
		}
	}

	d := make([][][]int, 2)
	for t := 0; t < 2; t++ {
		x0, y0 := 0, 0
		if t == 0 {
			x0, y0 = sx, sy
		} else {
			x0, y0 = gx, gy
		}

		d[t] = make([][]int, h)
		for i := range d[t] {
			d[t][i] = make([]int, w)
			for j := range d[t][i] {
				d[t][i][j] = INF
			}
		}

		d[t][x0][y0] = 0
		Q := make([][2]int, 0)
		Q = append(Q, [2]int{x0, y0})

		for len(Q) > 0 {
			x, y := Q[0][0], Q[0][1]
			Q = Q[1:]

			for k := 0; k < 4; k++ {
				x2, y2 := x+dx[k], y+dy[k]
				if 0 <= x2 && x2 < h && 0 <= y2 && y2 < w && B[x2][y2] != '#' && d[t][x2][y2] == INF {
					d[t][x2][y2] = d[t][x][y] + 1
					Q = append(Q, [2]int{x2, y2})
				}
			}
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if B[i][j] == '#' {
				res1, res2 := INF, INF
				for k := 0; k < 4; k++ {
					x, y := i+dx[k], j+dy[k]
					if 0 <= x && x < h && 0 <= y && y < w && B[x][y] != '#' {
						res1 = min(res1, d[0][x][y])
						res2 = min(res2, d[1][x][y])
					}
				}
				if res1 < INF && res2 < INF {
					ans = max(ans, res1+res2+2)
				}
			}
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
