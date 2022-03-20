package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	e = make([][]int, 3001)
	a = make([]int, 3001)
)

func dfs(x int) bool {
	for _, i := range e[x] {
		if a[x] > a[i] && !dfs(i) {
			return true
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}

	for i := 1; i <= n; i++ {
		if dfs(i) {
			fmt.Fprint(out, i, " ")
		}
	}
}
