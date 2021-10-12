package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	prime := PrimeFactorization(m)

	mod := int64(1e9 + 7)
	res := int64(1)
	for _, v := range prime {
		res *= nCrMod(int64(n+v-1), int64(v), mod)
		res %= mod
	}
	fmt.Println(res)
}

func PrimeFactorization(n int) map[int]int {
	res := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			res[i]++
			n /= i
		}
	}
	if n != 1 {
		res[n]++
	}
	return res
}

func powMod(a, n, mod int64) int64 {
	res := int64(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a, mod int64) int64 {
	return powMod(a, mod-2, mod)
}

func nCrMod(n, r, mod int64) int64 {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	res := int64(1)
	for i := int64(0); i < r; i++ {
		res = res * (n - i) % mod
		res = res * invMod(i+1, mod) % mod
	}
	return res
}
