package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 2023

	var b [N]int
	var s [N]string
	var f [N][N]int
	var qx, qy [N * N]int
	var sx, sy, gx, gy int

	dx := []int{0, 1, -1, 0, 0}
	dy := []int{0, 0, 0, 1, -1}

	var n, m int
	fmt.Fscan(in, &n, &m)
	b['>'] = 3
	b['v'] = 1
	b['<'] = 4
	b['^'] = 2
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		s[i] = " " + s[i]
		for j := 1; j <= m; j++ {
			if s[i][j] == 'S' {
				sx, sy = i, j
			} else if s[i][j] == 'G' {
				gx, gy = i, j
			} else if s[i][j] != '.' {
				f[i][j] = -1
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k, d := 1, b[s[i][j]]; i+k*dx[d] > 0 && j+k*dy[d] > 0 && d > 0; k++ {
				x := i + k*dx[d]
				y := j + k*dy[d]
				if x != 0 && y != 0 && x <= n && y <= m && f[x][y] != -1 {
					f[x][y] = -2
				} else {
					break
				}
			}
		}
	}
	if f[sx][sy] != 0 {
		fmt.Println(-1)
		return
	}
	f[sx][sy] = 1
	l, r := 1, 1
	qx[1] = sx
	qy[1] = sy
	for l <= r {
		x, y := qx[l], qy[l]
		l++
		for i := 1; i <= 4; i++ {
			xx := x + dx[i]
			yy := y + dy[i]
			if xx != 0 && yy != 0 && xx <= n && yy <= m && f[xx][yy] == 0 {
				r++
				qx[r] = xx
				qy[r] = yy
				f[xx][yy] = f[x][y] + 1
			}
		}
	}
	fmt.Println(f[gx][gy] - 1)
}
