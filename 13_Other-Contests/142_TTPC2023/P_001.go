package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	initMod()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	p := make([]int, n+1)
	p[0] = 1
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			p[j+1] = (p[j+1] + p[j]*a[i]%MOD) % MOD
		}
	}

	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = powMod(2, i*(i-1)/2)
		for j := 1; j < i; j++ {
			s[i] = (s[i] - (nCrMod(i-1, j-1)*s[j]%MOD)*powMod(2, (i-j)*(i-j-1)/2)%MOD + MOD) % MOD
		}
	}
	dp2 := make([][]int, n+1)
	for i := range dp2 {
		dp2[i] = make([]int, n+1)
	}
	r := make([]int, n+1)
	dp2[0][0] = 1
	dp2[1][1] = 1
	r[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 2; j < i+1; j++ {
			for k := 1; k < i+1; k++ {
				dp2[i][j] = (dp2[i][j] + ((nCrMod(i-1, k-1)*dp2[i-k][j-1]%MOD)*k%MOD)*r[k]%MOD) % MOD
			}
		}

		r[i] = s[i]
		pow_i := 1
		for j := 2; j < i+1; j++ {
			r[i] = (r[i] - dp2[i][j]*pow_i%MOD + MOD) % MOD
			pow_i = pow_i * i % MOD
		}
		dp2[i][1] = i * r[i] % MOD
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	q := make([]int, n+1)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			for k := 1; k <= i-j+1; k++ {
				dp[i][j] = (dp[i][j] + ((nCrMod(i-j, k-1)*dp[i-k][j-1]%MOD)*k%MOD)*r[k]%MOD) % MOD
			}
		}
	}

	q[1] = divMod(dp[n][1], n)
	pow_n := 1
	for i := 2; i < n+1; i++ {
		q[i] = dp[n][i] * pow_n % MOD
		pow_n = pow_n * n % MOD
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = (ans + p[i]*q[i]%MOD) % MOD
	}
	fmt.Println(ans)
}

const MOD = 998244353
const SIZE = 101010

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
