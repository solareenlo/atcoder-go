package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &k, &m)

	ans := n - m + 1
	a := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
	}

	initMod()

	for i := 1; i <= n-m; i++ {
		ans = ans * k % mod
	}

	const N = 25025
	const K = 404
	dp := [N][K]int{}
	sum := [N][K]int{}
	dp[1][1] = k
	sum[1][1] = k

	for i := 2; i <= n; i++ {
		for j := 1; j < k; j++ {
			dp[i][j] = (sum[i-1][j] + dp[i-1][j-1]*(k-j+1)) % mod
		}
		for j := k - 1; j > 0; j-- {
			sum[i][j] = (sum[i][j+1] + dp[i][j]) % mod
		}
	}

	l := 0
	lst := [K]int{}
	for i := 1; i <= m; i++ {
		l = max(l, lst[a[i]]+1)
		lst[a[i]] = i
		if i-l+1 >= k {
			fmt.Println(ans)
			return
		}
	}

	var tmp int
	if l != 1 {
		for i := 1; i <= k; i++ {
			lst[i] = m + 1
		}
		r := m
		for i := m; i > 0; i-- {
			r = min(r, lst[a[i]]-1)
			lst[a[i]] = i
		}
		for i := 1; i <= n-m+1; i++ {
			tmp = (tmp + sum[i+r-1][r]*sum[n-i-l+2][m-l+1]%mod) % mod
		}
		tmp = tmp * invf[r] % mod * invf[m-l+1] % mod
	} else {
		for i := 1; i <= n-m+1; i++ {
			for j := m; j < min(i+m, k); j++ {
				tmp = (tmp + dp[i+m-1][j]*sum[n-i-m+j+1][j]%mod*invf[j]%mod) % mod
			}
		}
		tmp = tmp * invf[m] % mod
	}

	fmt.Println((ans - tmp + mod) % mod)
}

const mod = 1000000007
const size = 101010

var fact, invf [size]int
var k int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * (k - i + 1)) % mod
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
