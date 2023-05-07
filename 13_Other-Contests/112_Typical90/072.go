package main

import (
	"bufio"
	"fmt"
	"os"
)

var H, W int
var S [16]string
var d [5]int = [5]int{0, 1, 0, -1}
var ex [16][16]bool
var ans int = -1

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &H, &W)
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &S[i])
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == '.' {
				dfs(i, j, i, j, 1)
			}
		}
	}
	fmt.Println(ans)
}

func dfs(x, y, s, t, c int) {
	ex[x][y] = true
	for r := 0; r < 4; r++ {
		tx := x + d[r]
		ty := y + d[r+1]
		if tx < 0 || ty < 0 || tx >= H || ty >= W || S[tx][ty] == '#' {
			continue
		}
		if ex[tx][ty] {
			if tx == s && ty == t && c > 2 && ans < c {
				ans = c
			}
		} else {
			dfs(tx, ty, s, t, c+1)
		}
	}
	ex[x][y] = false
}
