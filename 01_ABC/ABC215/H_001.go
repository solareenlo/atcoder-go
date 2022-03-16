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

	const N = 20
	a := make([]int, n)
	sa := make([]int, 1<<N)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sa[1<<i] = a[i]
	}
	b := make([]int, m)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &c[i][j])
		}
	}

	sb := make([]int, 1<<N)
	for j := 0; j < m; j++ {
		s := 0
		for i := 0; i < n; i++ {
			if c[i][j] != 0 {
				s |= (1 << i)
			}
		}
		sb[s] += b[j]
	}

	for j := 0; j < n; j++ {
		for i := 0; i < (1 << n); i++ {
			if (i>>j)&1 != 0 {
				sa[i] += sa[i^(1<<j)]
				sb[i] += sb[i^(1<<j)]
			}
		}
	}

	mn := 1 << 60
	for i := 1; i < (1 << n); i++ {
		if sa[i] < sb[i] {
			fmt.Println(0, 1)
			return
		}
		if sb[i] != 0 {
			mn = min(mn, sa[i]-sb[i])
		}
	}

	v := make([]int, 1<<N)
	for i := 0; i < (1 << n); i++ {
		if sa[i]-sb[i] == mn {
			v[i] = 1
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < (1 << n); i++ {
			if (i>>j&1) != 0 && v[i] >= 1 {
				v[i^(1<<j)] = 2
			}
		}
	}

	ans := 0
	for i := 0; i < (1 << n); i++ {
		if sa[i]-sb[i] == mn && v[i] == 1 {
			ans = (ans + nCrMod(sa[i], mn+1)) % mod
		}
	}
	fmt.Println(mn+1, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mod = 998244353
const size = 2000005

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
