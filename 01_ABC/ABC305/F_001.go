package main

import (
	"bufio"
	"fmt"
	"os"
)

var out = bufio.NewWriter(os.Stdout)

var st [101]bool
var g [101][]int

func get(x int) []int {
	var k int
	fmt.Scan(&k)
	pos := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&pos[i])
	}
	return pos
}

func dfs(x int) {
	st[x] = true
	g[x] = get(x)
	for _, y := range g[x] {
		if !st[y] {
			fmt.Fprintln(out, y)
			out.Flush()
			dfs(y)
			fmt.Fprintln(out, x)
			out.Flush()
			get(x)
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	dfs(1)
}
