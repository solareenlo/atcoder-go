package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 222222

var (
	e   = make([][]int, N)
	vis = make([]bool, N)
)

func dfs(x, fa int) {
	for _, y := range e[x] {
		if y != fa {
			dfs(y, x)
		}
	}
	if !vis[x] {
		if vis[fa] {
			fmt.Println("First")
			os.Exit(0)
		}
		vis[x] = true
		vis[fa] = true
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}
	vis[0] = true
	dfs(1, 0)
	fmt.Println("Second")
}
