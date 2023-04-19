package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

var n int
var a, b, ad, del [N]int
var e [N][]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}
	dfs(1, 0)
	for i := 0; i <= n; i++ {
		b[i] = 0
	}
	work(1, 0)
	solve(1, 0)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, a[i])
	}
}

func dfs(u, f int) {
	p := sum(f)
	add(u, 1)
	q := sum(u)
	for _, v := range e[u] {
		if v != f {
			dfs(v, u)
		}
	}
	ad[u] = u - 1 - (sum(u) - q)
	del[u] = sum(f) - p
}

func sum(x int) int {
	s := 0
	for i := x; i > 0; i -= i & (-i) {
		s += b[i]
	}
	return s
}

func add(x, v int) {
	for i := x; i <= n; i += i & (-i) {
		b[i] += v
	}
}

func work(u, f int) {
	a[1] += sum(n) - sum(u)
	add(u, 1)
	for _, v := range e[u] {
		if v != f {
			work(v, u)
		}
	}
	add(u, -1)
}

func solve(u, f int) {
	for _, v := range e[u] {
		if v != f {
			a[v] = a[u] + ad[v] - del[v]
			solve(v, u)
		}
	}
}
