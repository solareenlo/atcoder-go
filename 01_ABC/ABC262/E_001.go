package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	initMod()

	xy := make([]int, size)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		xy[u]++
		xy[v]++
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		if xy[i]%2 != 0 {
			cnt++
		}
	}

	ans := 0
	for i := 0; i <= k; i += 2 {
		ans += nCrMod(cnt, i) * nCrMod(n-cnt, k-i)
		ans %= mod
	}
	fmt.Println(ans)
}

const mod = 998244353

const size = 200005

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, mod-2)
}

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
