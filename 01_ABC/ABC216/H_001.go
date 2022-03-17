package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var K, N int
	fmt.Fscan(in, &K, &N)
	X := make([]int, K)
	for i := range X {
		fmt.Fscan(in, &X[i])
	}

	C := make([]int, 1010)
	C[0] = 1
	for i := 0; i < N; i++ {
		for j := i; j >= 0; j-- {
			C[j] = divMod(C[j], 2)
			C[j+1] += C[j]
			C[j+1] %= mod
		}
	}

	dp := [1 << 10][2020]int{}
	dp[0][0] = 1
	for i := 0; i < 1<<K; i++ {
		for j := 0; j < 2019; j++ {
			dp[i][j+1] += dp[i][j]
			dp[i][j+1] %= mod
			for l := 0; l < K; l++ {
				if (i >> l & 1) == 0 {
					t := j - X[l]
					if t >= 0 && t <= N {
						if bits.OnesCount(uint(i>>l))%2 != 0 {
							dp[i|1<<l][j+1] -= dp[i][j] * C[t] % mod
							dp[i|1<<l][j+1] += mod
							dp[i|1<<l][j+1] %= mod
						} else {
							dp[i|1<<l][j+1] += dp[i][j] * C[t] % mod
							dp[i|1<<l][j+1] %= mod
						}
					}
				}
			}
		}
	}

	fmt.Println(dp[(1<<K)-1][2019])
}

const mod = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}
