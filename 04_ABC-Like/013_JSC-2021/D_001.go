package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	fmt.Println(powMod(p-2, n-1) * (p - 1) % mod)
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
