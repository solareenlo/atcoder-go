package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	type node struct{ x, y int }
	a := make([]node, n+1)
	ds := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y)
		ds[i] = a[i].x
	}
	sort.Ints(ds)
	sort.Slice(a, func(i, j int) bool {
		return a[i].y < a[j].y
	})

	s := make([]int, n+1)
	s[0] = 1
	ps := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ps[lowerBound(ds[1:n+1], a[i].x)+1] = i
	}

	L := make([]int, n+2)
	L[n+1] = int(1e9)
	for i := n; i > 0; i-- {
		L[i] = min(L[i+1], ps[i])
	}

	R := make([]int, n+2)
	for i := 1; i <= n; i++ {
		R[i] = max(R[i-1], ps[i])
	}

	dp := make([]int, n+1)
	dp[0] = 1
	const MOD = 1_000_000_007
	for i, t := 1, 1; i <= n; i++ {
		for t <= n && R[t] <= i {
			tmp := 0
			if L[t] >= 2 {
				tmp = s[L[t]-2]
			}
			dp[i] = (2*dp[i] + s[i-1] - tmp) % MOD
			t++
		}
		s[i] = (s[i-1] + dp[i]) % MOD
	}

	fmt.Println((dp[n] + MOD) % MOD)
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
