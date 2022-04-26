package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, gx, gy int
	fmt.Fscan(in, &n, &gx, &gy)
	gx += 250
	gy += 250
	B := [500][500]bool{}
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		B[y+250][x+250] = true
	}

	const INF = 1 << 30
	d := [500][500]int{}
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			d[i][j] = INF
		}
	}
	d[250][250] = 0
	type pair struct{ y, x int }
	Q := make([]pair, 0)
	Q = append(Q, pair{250, 250})
	dx := []int{1, 1, 0, -1, -1, 0}
	dy := []int{0, 1, 1, 1, 0, -1}
	for len(Q) > 0 {
		y := Q[0].y
		x := Q[0].x
		Q = Q[1:]
		for k := 0; k < 6; k++ {
			y2 := y + dy[k]
			x2 := x + dx[k]
			if 0 <= y2 && y2 < 500 && 0 <= x2 && x2 < 500 && !B[y2][x2] && d[y2][x2] == INF {
				d[y2][x2] = d[y][x] + 1
				Q = append(Q, pair{y2, x2})
			}
		}
	}

	if d[gy][gx] < INF {
		fmt.Println(d[gy][gx])
	} else {
		fmt.Println(-1)
	}
}
