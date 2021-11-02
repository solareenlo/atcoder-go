package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	mod := 998244353
	dp := [3003]int{}
	res := 0
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		dp[0]++
		for j := s - v; j >= 0; j-- {
			dp[j+v] = (dp[j+v] + dp[j]) % mod
		}
		res = (res + dp[s]) % mod
	}

	fmt.Println(res)
}
