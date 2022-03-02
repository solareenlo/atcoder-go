package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)
	a--
	b--

	const o = 30000030
	const MOD = 998244353

	inv := make([]int, o)
	inv[1] = 1
	for i := 2; i <= a+2*b || i <= a*2; i++ {
		inv[i] = MOD - MOD/i*inv[MOD%i]%MOD
	}

	fac := make([]int, o)
	fac[0] = 1
	inv[0] = 1
	for i := 1; i <= a+2*b || i <= a*2; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = inv[i-1] * inv[i] % MOD
	}

	f := make([]int, b+1)
	for i := 0; i <= b; i++ {
		f[i] = fac[a+2*i] * inv[a] % MOD * inv[i] % MOD * inv[i] % MOD
	}

	ans := 0
	for i := 0; i <= b; i++ {
		ans = (ans + f[i]*f[b-i]) % MOD
	}
	for i := 0; i < b; i++ {
		ans = (ans - f[i]*4*f[b-i-1]) % MOD
	}

	fmt.Println(fac[2*a] * inv[a] % MOD * inv[a] % MOD * (ans + MOD) % MOD)
}
