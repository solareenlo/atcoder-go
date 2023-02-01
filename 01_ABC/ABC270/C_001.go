package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var a [][]int
var b []int
var e int

func main() {
	defer out.Flush()

	const N = 200020

	var n, s int
	fmt.Fscan(in, &n, &s, &e)

	a = make([][]int, N)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		a[u] = append(a[u], v)
		a[v] = append(a[v], u)
	}
	dfs(s, 0)
}

func dfs(x, px int) {
	b = append(b, x)
	if x == e {
		for i := range b {
			fmt.Fprintf(out, "%d ", b[i])
		}
		return
	}
	for i := range a[x] {
		if a[x][i] == px {
			continue
		}
		dfs(a[x][i], x)
	}
	b = b[:len(b)-1]
}
