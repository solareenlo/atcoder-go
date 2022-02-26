package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	A := make([]int, n+1)
	A[0] = 1 << 60
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &A[i])
	}

	const mod = 998244353
	dp := make([]int, n+1)
	dp[0] = 1
	tot := 0
	top := 0
	st := make([]int, n+1)
	D := make([]int, n+1)
	for i := 1; i <= n; i++ {
		R := A[i] * dp[i-1] % mod
		K := dp[i-1]
		for top != 0 {
			if A[st[top]] > A[i] {
				K += D[top]
				K %= mod
				R += D[top] * A[i] % mod
				R %= mod
				tot += mod - D[top]*A[st[top]]%mod
				tot %= mod
				top--
			} else {
				break
			}
		}
		tot += R
		tot %= mod
		dp[i] = (mod - tot) % mod
		top++
		st[top] = i
		D[top] = K
	}
	if n&1 != 0 {
		dp[n] = (mod - dp[n]) % mod
	}
	fmt.Println(dp[n])
}
