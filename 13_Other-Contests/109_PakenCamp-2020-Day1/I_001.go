package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type tuple struct {
		x, y, z int
	}

	var h, w, sx, sy, gx, gy int
	fmt.Fscan(in, &h, &w, &sx, &sy, &gx, &gy)
	sx--
	sy--
	gx--
	gy--

	var S [2000]string
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &S[i])
	}

	var D [2][2000][2000]int
	for i := range D {
		for j := range D[i] {
			for k := range D[i][j] {
				D[i][j][k] = -1
			}
		}
	}
	D[0][sx][sy] = 0
	D[1][sx][sy] = 0
	Q := make([]tuple, 0)
	Q = append(Q, tuple{0, sx, sy})
	Q = append(Q, tuple{1, sx, sy})
	for len(Q) > 0 {
		dir := Q[0].x
		x := Q[0].y
		y := Q[0].z
		Q = Q[1:]
		for k := 0; k < 4; k++ {
			nx := x + (k-1)%2
			ny := y + (k-2)%2
			if nx < 0 || nx >= h || ny < 0 || ny >= w {
				continue
			}
			if S[nx][ny] == '#' || dir != k%2 || D[dir^1][nx][ny] != -1 {
				continue
			}
			D[dir^1][nx][ny] = D[dir][x][y] + 1
			Q = append(Q, tuple{dir ^ 1, nx, ny})
		}
	}

	ans := -1
	for dir := 0; dir < 2; dir++ {
		if D[dir][gx][gy] == -1 {
			continue
		}
		if ans == -1 || ans > D[dir][gx][gy] {
			ans = D[dir][gx][gy]
		}
	}
	fmt.Println(ans)
}
