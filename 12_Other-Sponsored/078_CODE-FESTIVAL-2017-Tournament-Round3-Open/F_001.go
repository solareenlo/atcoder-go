package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007
	const N = 305

	var n int
	fmt.Fscan(in, &n)
	deg := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &deg[i])
	}

	var fac, inv, finv [N]int
	fac[0] = 1
	fac[1] = 1
	inv[1] = 1
	finv[0] = 1
	finv[1] = 1
	for i := 2; i <= n; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = (MOD - MOD/i) * inv[MOD%i] % MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}

	crl := 1
	for i := 1; i <= n; i++ {
		if deg[i] == 2 {
			crl &= 1
		} else {
			crl &= 0
		}
		if i == n && crl != 0 {
			fmt.Println(fac[n-1] * inv[2] % MOD)
			return
		}
	}

	var f [N][N][2]int
	f[0][0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i+1; j++ {
			for k := 0; k < 2; k++ {
				if deg[i] >= 1 {
					f[i][j][k] = (f[i][j][k] + f[i-1][j][k]*finv[deg[i]-1]%MOD) % MOD
				}
				if deg[i] >= 2 && j >= 1 {
					f[i][j][k] = (f[i][j][k] + f[i-1][j-1][k]*finv[deg[i]-2]%MOD) % MOD
				}
				if deg[i] >= 3 && j >= 1 && k != 0 {
					f[i][j][1] = (f[i][j][1] + f[i-1][j-1][0]*finv[deg[i]-3]%MOD) % MOD
				}
			}
		}
	}
	ans := 0
	for i := 3; i <= n; i++ {
		if n-i-1 >= 0 {
			ans = (ans + ((f[n][i][1]*fac[n-i-1]%MOD)*fac[i-1]%MOD)*inv[2]%MOD) % MOD
		}
	}
	fmt.Println(ans)
}
