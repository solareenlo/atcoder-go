package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	mod = m * m
	fmt.Println(powMod(10, n) / m)
}

var mod int

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
