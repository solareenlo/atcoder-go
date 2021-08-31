package main

import "fmt"

var (
	sx, sy, gx, gy int
	h, w, t, m     int
	dx             [4]int      = [4]int{0, 1, 0, -1}
	dy             [4]int      = [4]int{1, 0, -1, 0}
	d              [10][10]int = [10][10]int{}
	s              [10]string
)

func dfs(x, y, z int) {
	d[x][y] = z
	for i := 0; i < 4; i++ {
		nx := int(x + dx[i])
		ny := int(y + dy[i])
		if nx >= 0 && ny >= 0 && nx < h && ny < w {
			c := int(1)
			if s[nx][ny] == '#' {
				c = m
			}
			if d[nx][ny] > z+c {
				dfs(nx, ny, z+c)
			}
		}
	}
}

func fill() {
	for i := range d {
		for j := range d[i] {
			d[i][j] = 2e9
		}
	}
}

func main() {
	fmt.Scan(&h, &w, &t)

	for i := int(0); i < h; i++ {
		fmt.Scan(&s[i])
	}
	for i := int(0); i < h; i++ {
		for j := int(0); j < w; j++ {
			if s[i][j] == 'S' {
				sx, sy = i, j
			} else if s[i][j] == 'G' {
				gx, gy = i, j
			}
		}
	}

	l := int(1)
	r := int(t + 1)
	for r-l > 1 {
		m = (l + r) / 2
		fill()
		dfs(sx, sy, 0)
		if d[gx][gy] <= t {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(l)
}
