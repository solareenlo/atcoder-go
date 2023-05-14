package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a, b int
	fmt.Fscan(in, &n, &b, &a)
	n--
	if (n+a+b)&1 != 0 {
		fmt.Println(0)
		return
	}

	var fac, ifac [40000005]int
	fac[0] = 1
	N := max(a, b) + n
	for i := 1; i <= N; i++ {
		fac[i] = fac[i-1] * i % MOD
	}
	ifac[N] = ksm(fac[N], MOD-2, 1)
	for i := N; i > 0; i-- {
		ifac[i-1] = ifac[i] * i % MOD
	}

	ans := 0
	for i := max(0, a-n); i <= b; i++ {
		if ((i+n-a)&1) != 0 || ((b-i)&1) != 0 {
			continue
		}
		t := fac[i+n] * ifac[i] % MOD * ifac[(i+n-a)>>1] % MOD * ifac[(b-i)>>1] % MOD
		if (((b - i) >> 1) & 1) != 0 {
			ans -= t
		} else {
			ans += t
		}
	}
	fmt.Println(ksm(2, MOD-1-(n+b-a)/2, (ans%MOD+MOD)%MOD))
}

func ksm(x, tp, s int) int {
	for tp > 0 {
		if (tp & 1) != 0 {
			s = x * s % MOD
		}
		x = x * x % MOD
		tp >>= 1
	}
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
