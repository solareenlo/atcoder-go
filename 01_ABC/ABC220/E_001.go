package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	res := 0
	b := powMod(2, d-1)
	for k := n; k > d/2; k-- {
		if d < k {
			res += b * (d + 3) % mod
			res %= mod
		} else {
			res += b * (2*k - d - 1 + mod) % mod
			res %= mod
		}
		b *= 2
		b %= mod
	}

	fmt.Println(res)
}

const mod = 998244353

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
