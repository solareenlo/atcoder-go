package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	res := 0
	r := 1
	for i := 0; i < min(k, n-1)+1; i++ {
		res += r
		res %= mod
		r = r * (n - i) % mod * (n - i - 1) % mod
		r = divMod(r, i+1)
		r = divMod(r, i+1)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
