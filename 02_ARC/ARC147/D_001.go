package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(quickpow(n, m) * quickpow(m, n-1) % d)
}

const d = 998244353

func quickpow(b, p int) int {
	w := 1
	for p != 0 {
		if p%2 == 1 {
			w = w * b % d
		}
		b = b * b % d
		p = p >> 1
	}
	return w
}
