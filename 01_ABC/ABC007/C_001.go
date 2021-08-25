package main

import "fmt"

var dx [4]int = [4]int{1, 0, -1, 0}
var dy [4]int = [4]int{0, 1, 0, -1}

type p struct{ x, y int }

func main() {
	var r, c, sy, sx, gy, gx int
	fmt.Scan(&r, &c, &sy, &sx, &gy, &gx)

	m := make([]string, r)
	for i := range m {
		fmt.Scan(&m[i])
	}

	d := [50][50]int{}
	for i := range d {
		for j := range d[i] {
			d[i][j] = -1
		}
	}
	d[sy-1][sx-1] = 0

	q := []p{p{x: sx - 1, y: sy - 1}}
	for len(q) > 0 {
		y, x := q[0].y, q[0].x
		q = q[1:]
		for i := 0; i < 4; i++ {
			ny, nx := y+dy[i], x+dx[i]
			if ny < 0 || ny >= r || nx < 0 || nx >= c {
				continue
			}
			if m[ny][nx] == '#' {
				continue
			}
			if d[ny][nx] == -1 {
				q = append(q, p{x: nx, y: ny})
				d[ny][nx] = d[y][x] + 1
			}
		}
	}
	fmt.Println(d[gy-1][gx-1])
}
