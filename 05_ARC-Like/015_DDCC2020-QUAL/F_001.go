package main

import "fmt"

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)

	nn := gcd(n, p)
	n /= nn
	mm := gcd(m, p)
	m /= mm
	d := gcd(n, m)
	ans := (powMod(2, d) + powMod(2, n) + powMod(2, m) - 3) % mod
	ans = powMod(ans, nn*mm)
	fmt.Println(ans)
}

const mod = 1000000007

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
