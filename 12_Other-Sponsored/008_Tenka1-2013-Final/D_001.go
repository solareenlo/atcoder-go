package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var inv [33]int
	var dp [33][33][909]int

	inv[1] = 1
	for i := 2; i <= 30; i++ {
		inv[i] = inv[mod%i] * (mod - mod/i) % mod
	}

	var n, d int
	fmt.Fscan(in, &n, &d)
	dp[0][0][0] = 1
	s := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for j := 0; j < i+1; j++ {
			for k := 0; k < s+1; k++ {
				dp[i+1][j][k] = (dp[i+1][j][k] + dp[i][j][k]) % mod
				val := 1
				for l := 0; l < a; l++ {
					dp[i+1][j+1][k+l] = (dp[i+1][j+1][k+l] + dp[i][j][k]*val) % mod
					val = val * (d - k - l) % mod * inv[l+1] % mod
				}
			}
		}
		s += a
	}
	ret := 0
	for i := 0; i < n+1; i++ {
		for j := 0; j <= d && j <= s; j++ {
			val := dp[n][i][j] * powMod(n-i, d-j) % mod
			if i%2 == 0 {
				ret = (ret + val) % mod
			} else {
				ret = (ret - val + mod) % mod
			}
		}
	}
	fmt.Println(ret)
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
