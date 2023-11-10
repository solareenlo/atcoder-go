package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var vis [SIZE]bool

	initMod()

	var n, K int
	fmt.Fscan(in, &n, &K)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		vis[a] = true
	}
	ans := 0
	for i, j := 0, 0; ; i++ {
		if !vis[i] {
			ans = (ans + nCrMod(i+K-j-1, K-j)) % MOD
			j++
			if (j) > K {
				break
			}
		}
	}
	fmt.Println(ans)
}

const MOD = 998244353
const SIZE = 401010

var fact, invf [SIZE]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := 1; i < SIZE; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := 1
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
	if n < r {
		return 0
	}
	if n == r {
		return 1
	}
	if n < 0 || r < 0 {
		return 0
	}
	return (fact[n] * invf[r] % MOD) * invf[n-r] % MOD
}
