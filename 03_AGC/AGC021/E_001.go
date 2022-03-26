package main

import "fmt"

const mod = 998244353

func USA(x int) int {
	if x == 1 {
		return 1
	}
	return mod - mod/x*USA(mod%x)%mod
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	ans := 0
	sum := 1
	for i := k; i >= n; i-- {
		ans += sum
		ans %= mod
		sum = sum * (i - 1) % mod * USA(k-i+1) % mod
	}
	fmt.Println(ans)
}
