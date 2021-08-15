package main

import "fmt"

var mod = int64(1e9 + 7)
var mid = int64(1e8)

func main() {
	var n int64
	var k, m int
	fmt.Scan(&n, &k, &m)
	if n < int64(k) {
		dp := make([][]int, 11)
		for i := range dp {
			dp[i] = make([]int, 11)
		}
		dp[0][0] = 1
		for i := 1; i <= int(n); i++ {
			for j := 0; j < k; j++ {
				for l := 0; l < i; l++ {
					dp[i][(j+l)%k] += dp[i-1][j]
				}
			}
		}
		fmt.Println(dp[n][m])
	} else if n >= mod {
		fmt.Println(0)
	} else {
		b := [11]int64{1, 927880474, 933245637, 668123525, 429277690, 733333339, 724464507, 957939114, 203191898, 586445753, 698611116}
		res := b[n/mid]
		for i := n/mid*mid + 1; i <= n; i++ {
			res *= i
			res %= mod
		}
		res *= powmod(int64(k), mod-2)
		res %= mod
		fmt.Println(res)
	}
}

func powmod(x, n int64) int64 {
	res := int64(1)
	a := x
	for n > 0 {
		if n%2 == 1 {
			res *= a
			res %= mod
		}
		a *= a
		a %= mod
		n /= 2
	}
	return res
}
