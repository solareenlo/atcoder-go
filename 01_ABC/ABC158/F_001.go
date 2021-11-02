package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	type pair struct{ x, d int }
	p := make([]pair, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].d)
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})

	dp := make([]int, N+1)
	dp[N] = 1
	n := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		c := i + 1
		for c < N && p[c].x < p[i].x+p[i].d {
			c = n[c]
		}
		n[i] = c
		dp[i] = (dp[i+1] + dp[n[i]]) % mod
	}

	fmt.Println(dp[0])
}
