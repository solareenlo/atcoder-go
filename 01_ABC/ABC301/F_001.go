package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var s string
	fmt.Scan(&s)
	dp := make([]int, 29)
	dp[0] = 1
	use := 0
	for _, c := range s {
		dp2 := make([]int, 29)
		if isLower(string(c)) {
			for i := 0; i < 27; i++ {
				dp2[i] = dp[i]
			}
			dp2[28] = dp[27] + dp[28]
			if dp2[28] >= MOD {
				dp2[28] -= MOD
			}
		} else if isUpper(string(c)) {
			if ((use >> (c - 'A')) & 1) != 0 {
				dp2[27] = accumulate(dp[:28]) % MOD
			} else {
				cnt := popcount(uint32(use))
				for i := cnt; i < 27; i++ {
					dp2[i+1] = (dp2[i+1] + (dp[i]*(26-i)%MOD)*invMod(26-cnt)%MOD) % MOD
					dp2[27] = (dp2[27] + (dp[i]*(i-cnt)%MOD)*invMod(26-cnt)%MOD) % MOD
				}
				dp2[27] = (dp2[27] + dp[27]) % MOD
				use |= (1 << (c - 'A'))
			}
		} else {
			for i := 0; i < 27; i++ {
				dp2[i] = (dp2[i] + 26*dp[i]%MOD) % MOD
				dp2[i+1] = (dp2[i+1] + dp[i]*(26-i)%MOD) % MOD
				dp2[27] = (dp2[27] + dp[i]*i%MOD) % MOD
			}
			dp2[27] = (dp2[27] + 26*dp[27]%MOD) % MOD
			dp2[28] = 26 * (dp[27] + dp[28]) % MOD
		}
		dp, dp2 = dp2, dp
	}
	fmt.Println(accumulate(dp) % MOD)
}

func isUpper(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

func accumulate(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}

const MOD = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, MOD-2)
}
