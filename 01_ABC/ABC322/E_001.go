package main

import "fmt"

func main() {
	const MOD = 1000000007

	var n, k, p int
	fmt.Scan(&n, &k, &p)
	p++
	s := 1
	for i := 1; i <= k; i++ {
		s = s * p
	}
	var dp [100005]int
	for i := 0; i < s; i++ {
		dp[i] = MOD * MOD
	}
	dp[0] = 0
	var a [16]int
	for i := 1; i <= n; i++ {
		var c int
		fmt.Scan(&c)
		for j := 1; j <= k; j++ {
			fmt.Scan(&a[j])
		}
		for x := s - 1; x >= 0; x-- {
			y := x
			z := 0
			sb := 1
			for j := 1; j <= k; j++ {
				z += sb * min(y%p+a[j], p-1)
				y = y / p
				sb *= p
			}
			dp[z] = min(dp[z], dp[x]+c)
		}
	}
	if dp[s-1] < (MOD * MOD) {
		fmt.Println(dp[s-1])
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
