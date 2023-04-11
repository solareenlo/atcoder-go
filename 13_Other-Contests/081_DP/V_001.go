package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200010

var n, mod int
var e [N][]int
var f, g, l, r [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &mod)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}
	dfs(1, 0)
	dp(1, 0)
	fmt.Println(f[1])
	for i := 2; i <= n; i++ {
		fmt.Println(g[i] * f[i] % mod)
	}
}

func dfs(u, fa int) {
	for _, v := range e[u] {
		if v != fa {
			dfs(v, u)
		}
	}
	f[u] = 1
	for _, v := range e[u] {
		l[v] = f[u]
		f[u] = f[u] * (f[v] + 1) % mod
	}
	e[u] = reverseOrderInt(e[u])
	f[u] = 1
	for _, v := range e[u] {
		r[v] = f[u]
		f[u] = f[u] * (f[v] + 1) % mod
	}
	e[u] = reverseOrderInt(e[u])
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func dp(u, fa int) {
	if fa == 0 {
		g[u] = 1
	} else {
		g[u] = ((l[u]*r[u]%mod)*g[fa] + 1) % mod
	}
	for _, v := range e[u] {
		if v != fa {
			dp(v, u)
		}
	}
}
