package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	X := make([]int, n+1)
	X[0] = powMod(2, n-1) * n % mod
	for x := 1; x <= n; x++ {
		X[x] = (X[x-1] + (powMod(2, 2*x-1)*(2*x+1))%mod) % mod
	}

	allcountinv := powMod(invMod(2), n)

	for i := 0; i < n+1; i++ {
		X[i] *= allcountinv
		X[i] %= mod
	}

	for i := 0; i < n; i++ {
		ans := X[max(0, min(i-1, n-i-2))]
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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

func invMod(a int) int {
	return powMod(a, mod-2)
}

const mod = 998244353
