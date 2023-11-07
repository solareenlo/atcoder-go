package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var s [505]string
var f [505][505]bool
var t string = "snuke"
var dx []int = []int{-1, 1, 0, 0}
var dy []int = []int{0, 0, -1, 1}

func dfs(x, y, k int) {
	if f[x][y] || s[x][y] != t[k%5] {
		return
	}
	f[x][y] = true
	for i := 0; i < 4; i++ {
		dfs(x+dx[i], y+dy[i], k+1)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		s[i] = " " + s[i] + " "
	}
	s[0] = strings.Repeat(" ", len(s[1]))
	s[n+1] = strings.Repeat(" ", len(s[1]))
	dfs(1, 1, 0)
	if f[n][m] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
