package main

import (
	"fmt"
	"math/bits"
)

func getnum(x int) int { return 63 - countLeadingZeros(uint64(x)) }

func getid(x, y int) int {
	t := getnum(x)
	for i := t - 1; i >= 0; i-- {
		if ((x >> i) & 1) == y {
			return i
		}
	}
	return -1
}

func main() {
	const MAXN = 62
	const MAXM = 22

	var dp [2][(1 << MAXM) + 5]int
	var R [(1 << MAXM) + 5][2][2]int

	var n int
	var S, T string
	fmt.Scan(&n, &S, &T)
	S = " " + S
	T = " " + T
	lim := 0
	dp[lim][1] = 1
	up := 14
	for i := 1; i < (1 << up); i++ {
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 1; k++ {
				id := getid(i, j)
				if id == -1 {
					R[i][j][k] = 0
				} else {
					R[i][j][k] = (((i & ((1 << id) - 1)) | (1 << id)) << 1) | k
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < (1 << up); j++ {
			dp[lim^1][j] = 0
		}
		for j := 0; j < (1 << up); j++ {
			if dp[lim][j] == 0 {
				continue
			}
			u := 0
			if S[i+1] == 'X' {
				u = 1
			}
			v := 0
			if T[i+1] == 'X' {
				v = 1
			}
			if (j<<1 | u) < (1 << up) {
				dp[lim^1][(j<<1)|u] = max(dp[lim^1][(j<<1)|u], dp[lim][j])
			}
			if (j<<1 | v) < (1 << up) {
				dp[lim^1][(j<<1)|v] = max(dp[lim^1][(j<<1)|v], dp[lim][j])
			}
			if u == v {
				dp[lim^1][1] = max(dp[lim^1][1], (dp[lim][j]<<1)|u)
			}
			if R[j][u][v] != 0 {
				dp[lim^1][R[j][u][v]] = max(dp[lim^1][R[j][u][v]], (dp[lim][j]<<1)|u)
			}
			if R[j][v][u] != 0 {
				dp[lim^1][R[j][v][u]] = max(dp[lim^1][R[j][v][u]], (dp[lim][j]<<1)|v)
			}
		}
		lim ^= 1
	}
	ans := 1
	for i := 0; i < (1 << up); i++ {
		ans = max(ans, dp[lim][i])
	}
	t := getnum(ans)
	for i := t - 1; i >= 0; i-- {
		if ((ans >> i) & 1) != 0 {
			fmt.Print("X")
		} else {
			fmt.Print("Y")
		}
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countLeadingZeros(x uint64) int {
	return bits.LeadingZeros64(x)
}
