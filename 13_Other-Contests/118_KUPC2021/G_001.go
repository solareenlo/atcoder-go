package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	var P, inv [1 << 17]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &P[i])
		P[i]--
		inv[P[i]] = i
	}

	var A, B [1 << 17]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &B[i])
	}

	var vis [1 << 17]bool
	var dp [1 << 17][2][2]int
	ans := 0
	for i := 0; i < N; i++ {
		if !vis[i] {
			now := make([]int, 0)
			u := i
			now = append(now, u)
			vis[u] = true
			u = inv[u]
			for !vis[u] {
				now = append(now, u)
				vis[u] = true
				u = inv[u]
			}
			if len(now) == 1 {
				continue
			}
			dp[0][0][0] = A[now[0]]
			dp[0][0][1] = 1e18
			dp[0][1][1] = B[now[0]]
			dp[0][1][0] = 1e18
			for j := 1; j < len(now); j++ {
				dp[j][0][0] = min(dp[j-1][0][0], dp[j-1][0][1]) + A[now[j]]
				dp[j][0][1] = min(dp[j-1][0][0]+B[now[j-1]], dp[j-1][0][1]) + B[now[j]]
				dp[j][1][0] = min(dp[j-1][1][0], dp[j-1][1][1]) + A[now[j]]
				dp[j][1][1] = min(dp[j-1][1][0]+B[now[j-1]], dp[j-1][1][1]) + B[now[j]]
			}
			j := len(now) - 1
			ans += min(dp[j][0][0], dp[j][0][1], dp[j][1][0]+B[now[j]], dp[j][1][1])
		}
	}
	fmt.Println(ans)
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
