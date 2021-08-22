package main

import "fmt"

var mod int = 1e9 + 7
var r, c, x, y, d, l int

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b%2 == 1 {
		return pow(a, b-1) * a % mod
	}
	c := pow(a, b/2)
	return c * c % mod
}

func comb(a, b int) int {
	n, m := 1, 1
	for i := a - b + 1; i <= a; i++ {
		n = n * i % mod
	}
	for i := 1; i <= b; i++ {
		m = m * i % mod
	}
	return n * pow(m, mod-2) % mod
}

func a(x, y int) int {
	if x <= 0 || y <= 0 {
		return 0
	}
	if x*y < d+l {
		return 0
	}
	return comb(x*y, d) * comb(x*y-d, l) % mod
}

func main() {
	fmt.Scan(&r, &c, &x, &y, &d, &l)
	fmt.Println((r - x + 1) * (c - y + 1) % mod * ((a(x, y) - (a(x, y-1)+a(x-1, y))*2%mod + (a(x-1, y-1)*4+a(x-2, y)+a(x, y-2))%mod - (a(x-2, y-1)+a(x-1, y-2))*2%mod + a(x-2, y-2) + mod*3) % mod) % mod)
}
