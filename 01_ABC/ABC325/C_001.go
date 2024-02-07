package main

import (
	"fmt"
	"strings"
)

var g [1010][]string

func dfs(x, y int) {
	g[x][y] = "."
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if g[i][j] == "#" {
				dfs(i, j)
			}
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	g[0] = strings.Split(strings.Repeat(" ", m+2), "")
	for i := 1; i <= n; i++ {
		var s string
		fmt.Scan(&s)
		s = " " + s + " "
		g[i] = strings.Split(s, "")
	}
	g[n+1] = strings.Split(strings.Repeat(" ", m+2), "")
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if g[i][j] == "#" {
				ans++
				dfs(i, j)
			}
		}
	}
	fmt.Println(ans)
}
