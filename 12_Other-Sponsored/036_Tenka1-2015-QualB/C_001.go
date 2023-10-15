package main

import "fmt"

func main() {
	const MOD = 1000000007

	var n int
	fmt.Scan(&n)
	if n < 9 {
		fmt.Println(0)
		return
	}
	ans := (n - 6) / 2 % MOD
	ans = ans * (ans - 1 + MOD) % MOD
	ans = (ans + ans%2*MOD) / 2
	if (n & 1) != 0 {
		ans = (ans + ((n-7)/2+1)/2) % MOD
	}
	fmt.Println(ans)
}
