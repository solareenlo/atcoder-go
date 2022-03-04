package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n   int
	m   int
	sum = [size]int{}
)

func calc(k int) int {
	ans := n * nCrMod(m, k) % mod
	if k >= n {
		ans = (ans - sum[k+2] + mod) % mod
	} else {
		ans = (ans - ((n-k)*nCrMod(m, k)+sum[m-k+2])%mod + mod) % mod
	}
	return ans * fact[k] % mod * fact[m-k] % mod
}

func main() {
	in := bufio.NewReader(os.Stdin)

	initMod()

	fmt.Fscan(in, &m)
	n = (m >> 1)
	a := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : m+1]
	sort.Ints(tmp)

	for i := m; i > 0; i-- {
		sum[i] = (sum[i+2] + nCrMod(m, i)) % mod
	}

	ans := 0
	for i := 1; i <= m; i++ {
		if (i == 1) || (a[i] != a[i-1]) {
			ans = (ans + (a[i]-a[i-1])*calc(m-i+1)) % mod
		}
	}
	fmt.Println(ans)
}

const mod = 998244353
const size = 1000100

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
