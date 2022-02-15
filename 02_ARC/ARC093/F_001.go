package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i])
	}

	const N = 100005
	const M = 1_000_000_007
	li := 1 << n
	fac := make([]int, N)
	inv := make([]int, N)
	fac[0] = 1
	inv[li] = 1
	for i := 1; i <= li; i++ {
		fac[i] = fac[i-1] * i % M
	}

	x := fac[li]

	for t := M - 2; t > 0; t >>= 1 {
		if t&1 != 0 {
			inv[li] = inv[li] * x % M
		}
		x = x * x % M
	}

	for i := li; i > 0; i-- {
		inv[i-1] = inv[i] * i % M
	}

	dp := make([]int, N)
	dp[0] = 1
	for i := m - 1; i >= 0; i-- {
		for s := li - a[i] + 1; s > 0; s-- {
			for j := 1; j < 1<<n; j <<= 1 {
				if s&j != 0 {
					dp[s] = (dp[s] - dp[s^j]*j%M*fac[li-a[i]-(s^j)]%M*inv[li-a[i]-s+1]) % M
				}
			}
		}
	}

	ans := 0
	for s := 0; s < li; s++ {
		ans = (ans + dp[s]*fac[li-s-1]) % M
	}

	fmt.Println((ans*li%M + M) % M)
}
