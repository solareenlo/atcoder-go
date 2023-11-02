package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var fac [2 << 20]int

	fac[0] = 1
	for i := 1; i < 2<<20; i++ {
		fac[i] = fac[i-1] * i % MOD
	}
	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var N int
		fmt.Fscan(in, &N)
		fmt.Fprintln(out, divMod(((N-1)*(N*N%MOD-3)%MOD)*fac[2*N-4]%MOD, (fac[N-1]*fac[N-1])%MOD))
	}
}

const MOD = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a int) int {
	b, u, v := MOD, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}
