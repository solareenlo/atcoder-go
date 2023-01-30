package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	const mod = 998244353
	fmt.Println((mod + n%mod) % mod)
}
