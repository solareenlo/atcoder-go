package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	initMod()

	var n, m int
	fmt.Fscan(in, &n, &m)
	t := make([]int, n)
	u := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i], &u[i])
	}
	for i := 0; i < n; i++ {
		u[i] ^= t[i]
	}
	ans := 0
	for i := 0; i < n; i++ {
		cnt := (nCrMod(n, i) * powMod(2, i) % MOD) * powMod(2, m*(n-i)) % MOD
		if (i & 1) != 0 {
			ans = (ans + MOD - cnt) % MOD
		} else {
			ans = (ans + cnt) % MOD
		}
	}
	bs := make([]int, 0)
	for i := 0; i < n; i++ {
		x := u[i]
		for j := 0; j < len(bs); j++ {
			x = min(x, x^bs[j])
		}
		if x != 0 {
			bs = append(bs, x)
		}
	}
	allxor := 0
	for i := 0; i < n; i++ {
		allxor ^= t[i]
	}
	for j := 0; j < len(bs); j++ {
		allxor = min(allxor, allxor^bs[j])
	}
	if allxor == 0 {
		ans = (ans + (powMod(MOD-1, n)*powMod(2, (n-len(bs)))%MOD)*powMod(2, m)%MOD) % MOD
	}
	fmt.Println(ans)
}

const MOD = 998244353
const size = 400010

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, MOD-2)
}

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % MOD * invf[n-r] % MOD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
