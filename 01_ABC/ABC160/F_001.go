package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200015
const mod = 1000000007
const Maxn = 200005

var (
	n   int
	sz  = [N]int{}
	pre = [N]int{}
	inv = [N]int{}
	f   = [N]int{}
	to  = make([][]int, N)
)

func dfs(x, fa int) {
	sz[x] = 1
	for i := 0; i < len(to[x]); i++ {
		y := to[x][i]
		if y == fa {
			continue
		}
		dfs(y, x)
		sz[x] += sz[y]
	}
	f[1] = f[1] * inv[sz[x]] % mod
}

func calc(x, fa int) {
	for i := 0; i < len(to[x]); i++ {
		y := to[x][i]
		if y == fa {
			continue
		}
		f[y] = f[x] * sz[y] % mod * inv[n-sz[y]] % mod
		calc(y, x)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	pre[0] = 1
	for i := 1; i <= Maxn; i++ {
		pre[i] = pre[i-1] * i % mod
	}
	inv[1] = 1
	for i := 2; i <= Maxn; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
	}

	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		to[x] = append(to[x], y)
		to[y] = append(to[y], x)
	}
	f[1] = pre[n]
	dfs(1, 0)
	calc(1, 0)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, f[i])
	}
}
