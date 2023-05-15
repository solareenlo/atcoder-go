package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscan(in, &n, &t)
	P := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &P[i].y, &P[i].x)
	}
	sortPair(P)

	var dp [10001]int
	for i := range dp {
		dp[i] = t
	}
	dp[0] = 0
	var M [1001]int
	A := make([]int, 1000)
	for i := 0; i < n; i++ {
		for j, k := t, 0; j >= 0; j-- {
			for k < n && P[k].x+j <= t {
				k++
			}
			M[max(i, k)] = max(M[max(i, k)], t-dp[j])
			if j >= P[i].x {
				dp[j] = min(dp[j], dp[j-P[i].x]+P[i].y)
			}
		}
	}
	ans := 0
	for i, k := n-1, 0; i >= 0; i-- {
		A[i] = P[i].y
		sort.Ints(A[i:n])
		for k = i; k < n && M[i] >= A[k]; k++ {
			M[i] -= A[k]
		}
		ans = max(ans, k-i)
	}
	fmt.Println(ans)
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
