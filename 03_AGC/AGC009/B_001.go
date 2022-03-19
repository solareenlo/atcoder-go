package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	f = make([]int, 100005)
	e = make([][]int, 100005)
)

func dfs(u int) {
	for i := range e[u] {
		dfs(e[u][i])
		e[u][i] = f[e[u][i]]
	}
	sort.Ints(e[u])
	for i, a := 0, len(e[u]); i < a; i++ {
		f[u] = max(f[u], a-i+e[u][i])
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 2; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		e[a] = append(e[a], i)
	}
	dfs(1)
	fmt.Println(f[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
