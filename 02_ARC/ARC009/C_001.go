package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	s := 1 - k%2*2
	t := 1 - k%2*2

	for i := k; i > 2; i-- {
		t = t * -i % mod
		s += t
		s %= mod
	}

	for i := 1; i <= k; i++ {
		s = s * ((n - i + 1) % mod) % mod * powMod(i, mod-2) % mod
	}
	fmt.Println((s + mod) % mod)
}

const mod = 1777777777

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
