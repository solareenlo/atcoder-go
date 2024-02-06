package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var f, vac, t [10005]int
	f[0] = 1
	vac[0] = 1
	vac[1] = 1

	var n, x int
	fmt.Fscan(in, &n, &x)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &t[i])
	}
	for i := 2; i <= n; i++ {
		vac[i] = (MOD - MOD/i) * vac[MOD%i] % MOD
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= n; j++ {
			if i >= t[j] {
				f[i] = (f[i] + f[i-t[j]]*vac[n]%MOD) % MOD
			}
		}
	}

	ans := 0
	for i := max(x-t[1]+1, 0); i <= x; i++ {
		ans = (ans + f[i]*vac[n]%MOD) % MOD
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
