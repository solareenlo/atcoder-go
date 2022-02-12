package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	cnt int
	k   int
	v   = make([][]int, N)
	sz  = make([]int, N)
)

func dfs(x, f int) {
	maxi := 0
	sz[x] = 1
	for _, y := range v[x] {
		if y^f != 0 {
			dfs(y, x)
			sz[x] += sz[y]
			maxi = max(maxi, sz[y])
		}
	}
	maxi = max(maxi, n-sz[x])
	if maxi <= n/2 {
		cnt++
		k = x
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}

	initMod()

	dfs(1, 0)
	if cnt > 1 {
		fmt.Println(fact[n/2] * fact[n/2] % mod)
		return
	}

	dfs(k, 0)

	f := make([]int, N)
	f[0] = 1
	tt := 0
	for _, y := range v[k] {
		tt += sz[y]
		for i := tt; i >= 1; i-- {
			for j := 1; j < min(sz[y], i)+1; j++ {
				z := sz[y]
				f[i] = f[i] + f[i-j]*nCrMod(z, j)%mod*fact[z]%mod*invf[z-j]%mod
				f[i] %= mod
			}
		}
	}

	ans := 0
	for i := 0; i < n+1; i++ {
		tmp := 1
		if i&1 != 0 {
			tmp = mod - 1
		}
		ans = ans + tmp*f[i]%mod*fact[n-i]
		ans %= mod
	}
	fmt.Println((ans + mod) % mod)
}

const mod = 1000000007
const N = 5555

var fact, invf [N]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < N; i++ {
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
