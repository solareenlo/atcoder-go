package main

import (
	"bufio"
	"fmt"
	"os"
)

var H, W int
var Grid [105][105]bool
var visited [105][105 * 105]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &H, &W)
	for i := 0; i < H; i++ {
		var c string
		fmt.Fscan(in, &c)
		for j := 0; j < W; j++ {
			Grid[i][j] = (c[j] == '.')
		}
	}
	dfs(0, 0)
	if visited[H-1][W*100-1] {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}

func dfs(h, w int) {
	if h >= H || w >= W*100 {
		return
	}
	if !Grid[h][w%W] || visited[h][w] {
		return
	}
	visited[h][w] = true
	dfs(h+1, w)
	dfs(h, w+1)
}
