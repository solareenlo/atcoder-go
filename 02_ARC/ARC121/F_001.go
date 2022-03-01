package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 5
const mod = 998244353

var (
	f = [N][2]int{}
	e = make([][]int, N)
)

func dfs(x, fa int) {
	f[x][0] = 1
	f[x][1] = 1
	for _, y := range e[x] {
		if y != fa {
			dfs(y, x)
			t0 := (2*f[x][0]*f[y][0] + f[x][0]*f[y][1] + f[x][1]*f[y][0]) % mod
			t1 := (1*f[x][1]*f[y][0] + f[x][1]*f[y][1]) % mod
			f[x][0] = t0
			f[x][1] = t1
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}

	dfs(1, 0)

	ans := 1
	for i := 1; i < n*2; i++ {
		ans = 2 * ans % mod
	}
	fmt.Println((ans - f[1][0] + mod) % mod)
}
