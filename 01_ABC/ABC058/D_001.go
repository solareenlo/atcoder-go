package main

import "fmt"

var mod int = int(1e9 + 7)

func f(n int) int {
	res := 0
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		res += (2*i + 1 - n) * a % mod
		res %= mod
	}
	return res
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(f(n) * f(m) % mod)
}
