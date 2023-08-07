package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}
	var dx [6]int = [6]int{-1, -1, 0, 0, 1, 1}
	var dy [2][6]int = [2][6]int{{-1, 0, -1, 1, -1, 0}, {0, 1, -1, 1, 0, 1}}

	var C, R int
	fmt.Fscan(in, &C, &R)
	var c [110]string
	var a, d [110][110]int
	var sx, sy, tx, ty int
	for i := 0; i < C; i++ {
		fmt.Fscan(in, &c[i])
		for j := 0; j < R; j++ {
			d[i][j] = int(1e9)
			if c[i][j] == 's' {
				sx = i
				sy = j
			} else if c[i][j] == 't' {
				tx = i
				ty = j
			} else {
				a[i][j] = int(c[i][j] - '0')
			}
		}
	}
	q := make([]pair, 0)
	q = append(q, pair{sx, sy})
	d[sx][sy] = 0
	for len(q) > 0 {
		x := q[0].x
		y := q[0].y
		q = q[1:]
		for i := 0; i < 6; i++ {
			nx := x + dx[i]
			ny := y + dy[x%2][i]
			if 0 <= nx && nx < C && 0 <= ny && ny < R && d[nx][ny] > d[x][y]+a[nx][ny] {
				q = append(q, pair{nx, ny})
				d[nx][ny] = d[x][y] + a[nx][ny]
			}
		}
	}
	fmt.Println(d[tx][ty])
}
