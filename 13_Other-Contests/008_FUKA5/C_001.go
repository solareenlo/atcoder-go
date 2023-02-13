package main

import (
	"bufio"
	"fmt"
	"os"
)

var w, h, p, cnt int
var vis [30][30]bool
var plc [30][30]int

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for {
		cnt = 0
		for i := range vis {
			for j := range vis[i] {
				vis[i][j] = false
			}
		}
		fmt.Fscan(in, &w, &h, &p)
		if w == 0 && h == 0 && p == 0 {
			break
		}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				fmt.Fscan(in, &plc[i][j])
			}
		}
		for i := 0; i < p; i++ {
			var x, y int
			fmt.Fscan(in, &x, &y)
			dfs(y, x)
		}
		fmt.Fprintln(out, cnt)
	}
}

func dfs(x, y int) {
	if vis[x][y] {
		return
	}
	vis[x][y] = true
	cnt++
	for i := 0; i < 4; i++ {
		xx := x + dx[i]
		yy := y + dy[i]
		if xx >= 0 && xx < h && yy >= 0 && yy < w && plc[x][y] > plc[xx][yy] && vis[xx][yy] == false {
			dfs(xx, yy)
		}
	}
	return
}
