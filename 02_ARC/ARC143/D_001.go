package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 400004

var u, v, to, nx, head, bz, e []int
var tot int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	u = make([]int, N)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &u[i])
	}

	v = make([]int, N)
	to = make([]int, N)
	nx = make([]int, N)
	head = make([]int, N)
	tot = 1
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &v[i])
		add(u[i], v[i])
		add(v[i], u[i])
	}

	bz = make([]int, N)
	e = make([]int, N)
	for i := 1; i <= n; i++ {
		dfs(i)
	}
	for i := 1; i <= m; i++ {
		fmt.Fprintf(out, "%d", e[i<<1])
	}
}

func dfs(x int) {
	if bz[x] != 0 {
		return
	}
	bz[x] = 1
	for i := head[x]; i > 0; i = nx[i] {
		if e[i]+e[i^1] == 0 {
			e[i] = 1
			dfs(to[i])
		}
	}
}

func add(x, y int) {
	tot++
	to[tot] = y
	nx[tot] = head[x]
	head[x] = tot
}
