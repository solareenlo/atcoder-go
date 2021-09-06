package main

import "fmt"

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

func main() {
	var w, h int64
	fmt.Scan(&w, &h)
	fmt.Println(nCrMod(w+h-2, w-1, int64(1e9+7)))
}
