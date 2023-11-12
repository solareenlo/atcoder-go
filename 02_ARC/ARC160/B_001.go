package main

import "fmt"

func main() {
	const MOD = 998244353
	var solve func()
	solve = func() {
		var n int
		fmt.Scan(&n)
		ans := 0
		for x := 1; x*x <= n; x++ {
			ans = (ans + (n/x-x)*((x*6-3)%MOD) + (x*3 - 2)) % MOD
		}
		fmt.Println(ans)
	}

	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		solve()
	}
}
