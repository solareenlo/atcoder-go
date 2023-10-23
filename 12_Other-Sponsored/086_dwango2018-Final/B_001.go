package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MAX = 1000000
	const diff = 1 << 33
	const INF = 1 << 60

	var N, K int
	fmt.Fscan(in, &N, &K)
	dp := make([]int, MAX)
	dp[0] = -1
	for i := 1; i < MAX; i++ {
		dp[i] = INF
	}
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < N; i++ {
		for j := K; j >= 0; j-- {
			it := lowerBound(dp, A[i]+diff*j)
			dp[it] = A[i] + diff*j
		}
	}
	for i := 1; i < MAX; i++ {
		if dp[i] == INF {
			fmt.Fprintln(out, i-1)
			break
		}
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
