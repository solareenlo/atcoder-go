package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const inf int = 2000000000

var (
	N, T, K int
	A       [100100]int
	ma      [17][100100]int
	mi      [17][100100]int
	rm      [100100]int
	dp      [100100]int
)

func check(x int) bool {
	for i := 0; i < N; i++ {
		cur := i
		pre := T
		if i != 0 {
			pre = A[i-1]
		}
		for j := 16; j >= 0; j-- {
			if cur+(1<<j) > N {
				continue
			}
			if pre-x <= mi[j][cur] && ma[j][cur] <= pre+x {
				cur += 1 << j
			}
		}
		if cur > i {
			rm[i] = cur - 1
		} else {
			rm[i] = -1
		}
	}
	if rm[0] == -1 {
		return false
	}
	for i := 0; i <= N; i++ {
		dp[i] = inf
	}
	dp[0] = rm[0]
	for i := 1; i < N; i++ {
		if rm[i] == -1 {
			continue
		}
		k := sort.Search(N, func(j int) bool { return dp[j] >= i-1 })
		if dp[k] == inf {
			break
		}
		if dp[k+1] != inf {
			dp[k+1] = max(dp[k+1], rm[i])
		} else if dp[k] < rm[i] {
			dp[k+1] = rm[i]
		}
	}
	for i := 0; i <= K; i++ {
		if dp[i] == N-1 {
			return true
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &T, &K)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		ma[0][i] = A[i]
		mi[0][i] = A[i]
	}
	for i := 1; i < 17; i++ {
		for j := 0; j < N; j++ {
			if j+(1<<i) > N {
				continue
			}
			ma[i][j] = max(ma[i-1][j], ma[i-1][j+(1<<(i-1))])
			mi[i][j] = min(mi[i-1][j], mi[i-1][j+(1<<(i-1))])
		}
	}
	l := -1
	r := inf
	for r-l > 1 {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Println(r)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
