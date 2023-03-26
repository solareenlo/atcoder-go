package main

import (
	"fmt"
)

func pow_mod(a, n, m int) int {
	if n == 0 {
		return 1
	}
	x := pow_mod(a, n/2, m)
	ans := x * x % m
	if n%2 == 1 {
		ans = ans * a % m
	}
	return ans
}

func main() {
	var a, n, m int
	fmt.Scan(&a, &n, &m)
	fmt.Println(pow_mod(a, m, n))
}
