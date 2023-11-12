package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	fact := make([]int, n+1)
	for i := range fact {
		fact[i] = 1
	}
	for i := 1; i <= n; i++ {
		fact[i] = i * fact[i-1] % MOD
	}
	ans := 0
	for i := 0; i <= n-k; i++ {
		ans = (ans - ((divMod(fact[n-i-1], (fact[k-1]*fact[n-i-k])%MOD) % MOD) * a[i] % MOD) + MOD) % MOD
	}
	for i := k - 1; i < n; i++ {
		ans = (ans + (divMod(fact[i], (fact[k-1]*fact[i-k+1])%MOD) * a[i] % MOD)) % MOD
	}
	ans = ans * divMod((fact[k]*fact[n-k]%MOD), fact[n]) % MOD
	fmt.Println(ans)
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
