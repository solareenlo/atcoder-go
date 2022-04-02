package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	const mod = 998244353
	p := 10
	q := 1
	ans := 0
	for {
		m := min(n, p-1) - q + 1
		m %= mod
		ans = (ans + (m*(m+1)/2)%mod) % mod
		if p > n {
			break
		}
		p *= 10
		q *= 10
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
