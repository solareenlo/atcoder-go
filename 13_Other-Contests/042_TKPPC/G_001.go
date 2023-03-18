package main

import (
	"fmt"
)

const (
	p    = 1000000007
	INF  = 1000000000
	MAXN = 1000000
)

var (
	N     int
	str   string
	phash []uint64
	power []uint64
	dp    []int
)

func RollingHash() {
	phash = make([]uint64, N+1)
	power = make([]uint64, N+1)
	phash[0] = 0
	power[0] = 1

	for i := 0; i < N; i++ {
		phash[i+1] = uint64(str[i]-'0') + phash[i]*p
		power[i+1] = power[i] * p
	}
}

func H(l, r int) uint64 {
	return phash[r] - phash[l]*power[r-l]
}

func LCP(a, b, n int) int {
	lb := -1
	ub := n
	for ub-lb > 1 {
		mid := (ub + lb) / 2
		if H(a, a+mid+1) == H(b, b+mid+1) {
			lb = mid
		} else {
			ub = mid
		}
	}
	return lb
}

func main() {
	fmt.Scan(&N, &str)
	RollingHash()
	dp = make([]int, MAXN)
	for i := 0; i < MAXN; i++ {
		dp[i] = INF
	}
	dp[N-1] = 1
	for i := N - 1; i > 0; i-- {
		dp[i-1] = min(dp[i]+1, dp[i-1])
		if i-dp[i] < 0 {
			continue
		}
		if str[i] == '0' {
			continue
		}
		lcp := LCP(i-dp[i], i, dp[i])
		if lcp+1 == dp[i] || str[i-dp[i]+lcp+1] < str[i+lcp+1] {
			if i-dp[i]-1 < 0 {
				continue
			}
			dp[i-dp[i]-1] = min(dp[i-dp[i]-1], dp[i]+1)
		} else {
			dp[i-dp[i]] = min(dp[i-dp[i]], dp[i])
		}
	}
	fmt.Println(dp[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
