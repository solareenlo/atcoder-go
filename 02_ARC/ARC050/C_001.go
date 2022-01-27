package main

import "fmt"

var mod int

func solve(a, x int) int {
	if x == 1 {
		return 1
	}
	q := solve(a, x/2)
	if x&1 != 0 {
		return (((q*powMod(a, x/2)%mod+q)%mod)*a + 1) % mod
	} else {
		return (q*powMod(a, x/2)%mod + q) % mod
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b, &mod)

	g := gcd(a, b)
	fmt.Println(solve(10, a) * solve(powMod(10, g), b/g) % mod)
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
