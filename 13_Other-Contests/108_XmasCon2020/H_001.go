package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const U = 2097152
	const mod = 998244353
	const inv2 = (mod + 1) / 2

	var n int
	fmt.Fscan(in, &n)

	var s string
	fmt.Fscan(in, &s)
	var l [U]int
	for i := 0; i < (1 << n); i++ {
		l[i] = int(s[i]) ^ 48
	}
	fmt.Fscan(in, &s)
	var r [U]int
	for i := 0; i < (1 << n); i++ {
		r[i] = int(s[i]) ^ 48
	}

	var sz [U]int
	var L [U][22]int
	for S := 0; S < (1 << n); S++ {
		sz[S] = sz[S>>1] + (S & 1)
		L[S][sz[S]] = l[S]
	}
	for i := 0; i < n; i++ {
		for S := 0; S < (1 << n); S++ {
			if ((S >> i) & 1) != 0 {
				for j := 0; j <= n; j++ {
					L[S][j] = L[S][j] + L[S^(1<<i)][j] - mod
					L[S][j] += (L[S][j] >> 31) & mod
				}
			}
		}
	}

	var f [U][22]int
	var g [22]int
	for S := 0; S < (1 << n); S++ {
		g[0] = 1
		f[S][0] = 0
		for i := 1; i <= n; i++ {
			w := mod - 2*L[S][i]%mod
			for j := 1; j < i; j++ {
				w = w - g[j]*g[i-j]%mod
				w += (w >> 31) & mod
			}
			g[i] = w * inv2 % mod
			f[S][i] = -g[i]
			f[S][i] += (f[S][i] >> 31) & mod
		}
	}
	for i := 0; i < n; i++ {
		for S := 0; S < (1 << n); S++ {
			if ((S >> i) & 1) != 0 {
				for j := 0; j <= n; j++ {
					f[S][j] = f[S][j] - f[S^(1<<i)][j]
					f[S][j] += (f[S][j] >> 31) & mod
				}
			}
		}
	}

	ans := 0
	for S := 0; S < (1 << n); S++ {
		if r[S] != 0 {
			ans = ans + f[S][sz[S]] - mod
			ans += (ans >> 31) & mod
		}
	}
	fmt.Println(ans)
}
