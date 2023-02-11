package main

import "fmt"

func main() {
	const mod = 998244353

	var n int
	fmt.Scan(&n)

	ans := 2
	for i := 2 * n; i >= n+2; i-- {
		ans *= 2 * i
		ans %= mod
	}
	fmt.Println(ans)
}
