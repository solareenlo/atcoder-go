package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	vis = [110]bool{}
	g   = [110][110]int{}
	ans = [10010]int{}
	c   = [110]int{}
)

func DFS(x int) {
	vis[x] = true
	for y := 1; y <= n; y++ {
		if g[x][y] != 0 {
			i := abs(g[x][y])
			if ans[i] == 0 {
				if c[x] >= c[y] {
					ans[i] = g[x][y]
				} else {
					ans[i] = -g[x][y]
				}
			}
			if c[x] == c[y] && !vis[y] {
				DFS(y)
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m int
	fmt.Fscan(in, &n, &m)

	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		g[x][y] = i
		g[y][x] = -i
	}

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
	}

	for i := 1; i <= n; i++ {
		if !vis[i] {
			DFS(i)
		}
	}

	for i := 1; i <= m; i++ {
		if ans[i] > 0 {
			fmt.Fprintln(out, "->")
		} else {
			fmt.Fprintln(out, "<-")
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
