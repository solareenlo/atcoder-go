package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	res, comb := 0, 1
	for i := 0; i < n-k+1; i++ {
		res += comb * (a[k+i-1] - a[n-i-k] + mod) % mod
		res %= mod
		comb = divMod(comb*(k+i)%mod, (i + 1))
	}

	fmt.Println(res)
}

const mod = 1000000007

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
