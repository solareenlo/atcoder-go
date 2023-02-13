package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	ma := make([][]string, n+1)
	var sx, sy, cx, cy, gx, gy int
	for i := 1; i <= n; i++ {
		var t string
		fmt.Fscan(in, &t)
		t = "#" + t
		ma[i] = strings.Split(t, "")
		for j := 1; j <= m; j++ {
			if ma[i][j] == "S" {
				sx = i
				sy = j
			}
			if ma[i][j] == "C" {
				cx = i
				cy = j
			}
			if ma[i][j] == "G" {
				gx = i
				gy = j
			}
		}
	}

	type pair struct {
		x, y int
	}

	nx := [5]int{0, 1, -1, 0, 0}
	ny := [5]int{0, 0, 0, 1, -1}
	q := make([]pair, 0)
	q = append(q, pair{cx, cy})
	var flag [505][505]int
	for len(q) > 0 {
		dx := q[0].x
		dy := q[0].y
		q = q[1:]
		for i := 1; i <= 4; i++ {
			xx := dx + nx[i]
			yy := dy + ny[i]
			if xx < 1 || xx > n || yy < 1 || yy > m || flag[xx][yy] != 0 || ma[xx][yy] == "#" {
				continue
			}
			q = append(q, pair{xx, yy})
			flag[xx][yy] = flag[dx][dy] + 1
		}
	}

	var ans int
	if flag[sx][sy] != 0 && flag[gx][gy] != 0 {
		ans = flag[sx][sy] + flag[gx][gy]
	} else {
		ans = -1
	}
	fmt.Println(ans)
}
