package main

import "fmt"

func main() {
	const MOD = 1000000007

	var s string
	var k int
	fmt.Scan(&s, &k)
	n := len(s)
	st := make([][]int, 0)
	ans := 1
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			dp := make([]int, 3)
			dp[0] = 1
			dp[2] = 1
			dp[1] = 2
			st = append(st, dp)
		} else {
			dp := st[len(st)-1]
			st = st[:len(st)-1]
			a := len(dp) - 1
			for i := 0; i < a+1; i++ {
				u := a - i
				if abs(u-i) > k {
					dp[i] = 0
				}
			}
			if len(st) == 0 {
				add := 0
				for i := 0; i < a+1; i++ {
					add = (add + dp[i]) % MOD
				}
				ans = (ans * add) % MOD
				continue
			}
			par := st[len(st)-1]
			st = st[:len(st)-1]
			b := len(par) - 1
			next := make([]int, a+b+1)
			for i := 0; i < a+1; i++ {
				for j := 0; j < b+1; j++ {
					next[i+j] = (next[i+j] + dp[i]*par[j]) % MOD
				}
			}
			st = append(st, next)
		}
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
