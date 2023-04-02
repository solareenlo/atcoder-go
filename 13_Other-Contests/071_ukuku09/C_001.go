package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1e18)

var dx [5]int = [5]int{-1, 0, 1, 0, 0}
var dy [5]int = [5]int{0, -1, 0, 1, 0}
var mv string = "^<v>"
var s [40]string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w, Q int
	fmt.Fscan(in, &h, &w, &Q)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &s[i])
	}
	var d [40][40][40][40]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < h; k++ {
				for l := 0; l < w; l++ {
					d[i][j][k][l] = INF
				}
			}
		}
	}

	type pair struct {
		x, y int
	}
	q := make([]pair, 0)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				continue
			}
			q = append(q, pair{i*w + j, i*w + j})
			d[i][j][i][j] = 0
		}
	}
	for len(q) > 0 {
		xx := q[0].x
		yy := q[0].y
		q = q[1:]
		xi := xx / w
		yi := xx % w
		xj := yy / w
		yj := yy % w
		for i := 0; i < 5; i++ {
			nx := xi + dx[i]
			ny := yi + dy[i]
			if !can(nx, ny, xi, yi) {
				continue
			}
			for j := 0; j < 5; j++ {
				mx := xj + dx[j]
				my := yj + dy[j]
				if !can(mx, my, xj, yj) {
					continue
				}
				if d[nx][ny][mx][my] > d[xi][yi][xj][yj]+1 {
					d[nx][ny][mx][my] = d[xi][yi][xj][yj] + 1
					q = append(q, pair{nx*w + ny, mx*w + my})
				}
			}
		}
	}
	for Q > 0 {
		Q--
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &y1, &x1, &y2, &x2)
		x1--
		y1--
		x2--
		y2--
		res := d[x1][y1][x2][y2]
		if res == INF {
			res = -1
		}
		fmt.Fprintln(out, res)
	}
}

func can(x1, y1, x2, y2 int) bool {
	if s[x1][y1] == '#' {
		return false
	}
	if s[x1][y1] != '.' {
		for i := 0; i < 4; i++ {
			if mv[i] != s[x1][y1] {
				continue
			}
			nx := x1 + dx[i]
			ny := y1 + dy[i]
			if s[nx][ny] == '#' {
				nx = x1
				ny = y1
			}
			if nx == x2 && ny == y2 {
				return true
			}
		}
		return false
	}
	return true
}
