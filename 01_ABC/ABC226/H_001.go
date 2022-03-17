package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscan(in, &n, &t)

	l := make([]int, n)
	r := make([]int, n)
	ha := make([]int, 2*n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		ha[2*i] = l[i]
		ha[2*i+1] = r[i]
	}
	sort.Ints(ha)

	ans := 0
	for i := 0; i < 2*n-1; i++ {
		a := ha[i]
		b := ha[i+1]
		dp := [50 + 1][50 + 1][50 + 1]int{}
		dp[0][0][0] = 1
		for j := 0; j < n; j++ {
			in := 0
			ue := 0
			sh := 0
			if a < r[j] && l[j] < b {
				in = (min(b, r[j]) - max(a, l[j]) + mod) % mod
			}
			if b < r[j] {
				ue = (r[j] - max(b, l[j]) + mod) % mod
			}
			in = divMod(in, (r[j]-l[j]+mod)%mod)
			ue = divMod(ue, (r[j]-l[j]+mod)%mod)
			sh = (1 - in + mod - ue + mod) % mod
			for k := 0; k < j+1; k++ {
				for l := 0; l < j+1; l++ {
					w := dp[j][k][l]
					dp[j+1][k+1][l] += w * in % mod
					dp[j+1][k+1][l] %= mod
					dp[j+1][k][l+1] += w * ue % mod
					dp[j+1][k][l+1] %= mod
					dp[j+1][k][l] += w * sh % mod
					dp[j+1][k][l] %= mod
				}
			}
		}
		for j := 0; j < t; j++ {
			for k := t - j; k <= n; k++ {
				tmp1 := (b - a + mod) % mod
				tmp2 := (t - j + mod) % mod
				ans += dp[n][k][j] * ((b - divMod(tmp1*tmp2%mod, (k+1)) + mod) % mod) % mod
				ans %= mod
			}
		}
	}
	fmt.Println(ans)
}

const mod = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
