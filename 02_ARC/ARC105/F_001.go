package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	lim := (1 << n) - 1
	const N = 133333
	l := make([]int, N)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		S := lim ^ (1 << u) ^ (1 << v)
		for T := S; ; T = (T - 1) & S {
			l[lim^T]++
			if T == 0 {
				break
			}
		}
	}

	const mod = 998244353
	pw2 := make([]int, m+1)
	pw2[0] = 1
	for i := 1; i <= m; i++ {
		pw2[i] = (pw2[i-1] << 1) % mod
	}

	g := make([]int, lim+1)
	f := make([]int, lim+1)
	for S := 1; S <= lim; S++ {
		S2 := S ^ (S & -S)
		for T := S2; ; T = (T - 1) & S2 {
			g[S] = (g[S] + pw2[l[S]-l[T]-l[S^T]]) % mod
			f[S] = (f[S] + f[S^T]*g[T]) % mod
			if T == 0 {
				break
			}
		}
		g[S] = (g[S] << 1) % mod
		f[S] = (g[S] + mod - f[S]) % mod
	}

	fmt.Println(f[lim] * (mod + 1) / 2 % mod)
}
