package main

import (
	"bufio"
	"fmt"
	"os"
)

var H, W int
var c, p [11][11]int
var ret int
var dir [4]int = [4]int{0, 1, 0, -1}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &H, &W)
	sx, sy, sd := -1, -1, 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &c[i][j])
			p[i][j] = (i*W + j + 1) % (H * W)
			if c[i][j] == 0 {
				sx = i
				sy = j
			}
			if c[i][j] != p[i][j] {
				sd++
			}
		}
	}
	ret = 24
	solve(sx, sy, 0, -1, sd)
	fmt.Println(ret)
}

func solve(x, y, depth, dpre, diff int) {
	if diff == 0 {
		ret = depth
		return
	}
	if depth+diff > ret {
		return
	}
	for i := 0; i < 4; i++ {
		if (i ^ 2) == dpre {
			continue
		}
		tx := x + dir[i]
		ty := y + dir[i^1]
		if 0 <= tx && tx < H && 0 <= ty && ty < W {
			c[x][y], c[tx][ty] = c[tx][ty], c[x][y]
			tmp0, tmp1, tmp2, tmp3 := 0, 0, 0, 0
			if c[x][y] == p[x][y] {
				tmp0 = 1
			}
			if c[tx][ty] == p[tx][ty] {
				tmp1 = 1
			}
			if c[x][y] == p[tx][ty] {
				tmp2 = 1
			}
			if c[tx][ty] == p[x][y] {
				tmp3 = 1
			}
			solve(tx, ty, depth+1, i, diff-tmp0-tmp1+tmp2+tmp3)
			c[x][y], c[tx][ty] = c[tx][ty], c[x][y]
		}
	}
}
