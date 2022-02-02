package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	const N = 101010
	a := make([]int, N)
	m := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			m++
			a[m] = i
			if n/i != i {
				m++
				a[m] = n / i
			}
		}
	}
	sort.Ints(a[:m+1])

	f := make([]int, m+1)
	ans := 0
	for i := 1; i <= m; i++ {
		f[i] = powMod(k, (a[i]+1)/2)
		for j := 1; j < i; j++ {
			if a[i]%a[j] == 0 {
				f[i] = (f[i] - f[j] + mod) % mod
			}
		}
		if a[i]%2 != 0 {
			ans = (ans + f[i]*a[i]) % mod
		} else {
			ans = (ans + a[i]*powMod(2, mod-2)%mod*f[i]) % mod
		}
	}
	fmt.Println(ans)
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
