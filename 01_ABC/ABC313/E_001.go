package main

import "fmt"

func main() {
	const MOD = 998244353

	var n int
	var s string
	fmt.Scan(&n, &s)
	for i := 0; i < n; i++ {
		if s[i] != '1' {
			if i != n-1 && s[i+1] != '1' {
				fmt.Println(-1)
				return
			}
		}
	}
	ans := 0
	for i := n - 1; i >= 1; i-- {
		ans++
		ans += ans * (int(s[i]-'0') - 1)
		ans %= MOD
	}
	fmt.Println(ans)
}
