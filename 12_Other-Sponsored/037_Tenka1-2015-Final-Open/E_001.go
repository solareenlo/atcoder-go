package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var n int
	fmt.Fscan(in, &n)
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i][1], &a[i][0])
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] > a[j][1]
		}
		return a[i][0] > a[j][0]
	})

	inv := make([]int, n+1)
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = mod - (mod/i)*inv[mod%i]%mod
	}

	dp := make([]int, n+1)
	prd, sum1, sum2 := 2, 0, 0

	for i := 2; i < n; i++ {
		prd = (prd * (i + 1)) % mod
		ainv := (inv[i] * inv[i+1]) % mod
		sum1 = (sum1 + ainv) % mod
		sum2 = (sum2 + a[i-1][0]*ainv%mod) % mod
		tmp := (2 * prd * a[i][1]) % mod
		tmp = tmp * ((sum2 - a[i][0]*sum1%mod) + mod) % mod
		dp[i+1] = (dp[i]*(i+1)%mod + tmp) % mod
	}

	if dp[n] < 0 {
		dp[n] += mod
	}

	fmt.Println(dp[n])
}
