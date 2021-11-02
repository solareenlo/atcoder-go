package main

import "fmt"

const mod = 998244353

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	res := 0
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	dp := [3003]int{}
	for i := 0; i < n; i++ {
		dp[0]++
		for j := k; j >= a[i]; j-- {
			dp[j] = (dp[j] + dp[j-a[i]]) % mod
		}
		res = (res + dp[k]) % mod
	}

	fmt.Println(res)
}
