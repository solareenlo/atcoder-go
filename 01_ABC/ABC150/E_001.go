package main

import (
	"fmt"
	"sort"
)

const mod = 1000000007

func main() {
	var n int
	fmt.Scan(&n)

	c := make([]int, n)
	for i := range c {
		fmt.Scan(&c[i])
	}
	sort.Ints(c)

	res := 0
	for i := 0; i < n; i++ {
		res += c[i] * (n + 1 - i) % mod
		res %= mod
	}

	fmt.Println(res * powMod(4, n-1) % mod)
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
