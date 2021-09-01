package main

import "fmt"

func powMod(x, n, mod int64) int64 {
	res := int64(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		n /= 2
	}
	return res
}

func invMod(x, mod int64) int64 {
	return powMod(x, mod-2, mod)
}

func nCr(n, r, mod int64) int64 {
	res := int64(1)
	for i := int64(0); i < r; i++ {
		res = res * (n - i) % mod
		res = res * invMod(i+1, mod) % mod
	}
	return res
}

func main() {
	var n, k int64
	fmt.Scan(&n, &k)

	fmt.Println(nCr(n+k-1, k, int64(1e9+7)))
}
