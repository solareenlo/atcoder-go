package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var N, B, K int
	fmt.Fscan(in, &N, &B, &K)

	var dp [1000]int
	for i := 0; i < K; i++ {
		var c int
		fmt.Fscan(in, &c)
		dp[c%B] += 1
	}

	var ans, tmp [1000]int
	ans[0] = 1
	t := 10
	for N != 0 {
		if N&1 != 0 {
			for j := 0; j < B; j++ {
				tmp[j] = ans[j]
				ans[j] = 0
			}
			for j := 0; j < B; j++ {
				for l := 0; l < B; l++ {
					ans[(j*t+l)%B] = (ans[(j*t+l)%B] + tmp[j]*dp[l]%MOD) % MOD
				}
			}
		}
		N >>= 1
		for j := 0; j < B; j++ {
			tmp[j] = dp[j]
			dp[j] = 0
		}
		for j := 0; j < B; j++ {
			for l := 0; l < B; l++ {
				dp[(j*t+l)%B] = (dp[(j*t+l)%B] + tmp[j]*tmp[l]%MOD) % MOD
			}
		}
		t = t * t % B
	}
	fmt.Println(ans[0])
}
