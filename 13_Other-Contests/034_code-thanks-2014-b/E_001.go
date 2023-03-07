package main

import (
	"bufio"
	"fmt"
	"os"
)

var r, c int
var dx [4]int = [4]int{1, 0, -1, 0}
var dy [4]int = [4]int{0, 1, 0, -1}
var g [50][50]bool
var visited [50][50]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	var rs, cs, rg, cg, n int
	fmt.Fscan(in, &r, &c, &rs, &cs, &rg, &cg, &n)
	for i := 0; i < n; i++ {
		var r, c, h, w int
		fmt.Fscan(in, &r, &c, &h, &w)
		for j := r; j < r+h; j++ {
			for k := c; k < c+w; k++ {
				g[j][k] = true
			}
		}
	}
	dfs(rs, cs)
	if visited[rg][cg] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func dfs(x, y int) {
	if x < 1 || y < 1 || x > c+1 || y > r+1 || !g[y][x] || visited[y][x] {
		return
	}
	visited[y][x] = true
	for i := 0; i < 4; i++ {
		dfs(x+dx[i], y+dy[i])
	}
}
