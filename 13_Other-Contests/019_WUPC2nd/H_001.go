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

	var N, M int
	fmt.Fscan(in, &N, &M)

	l := make([]int, M)
	r := make([]int, M)
	t := make([]int, 0)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		l[i]--
		t = append(t, l[i])
		t = append(t, r[i])
	}
	t = append(t, N)
	t = append(t, 0)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	t = unique(t)
	type pair struct {
		x, y int
	}
	P := make([]pair, 0)
	for i := 0; i < M; i++ {
		l[i] = lowerBound(t, l[i])
		r[i] = lowerBound(t, r[i])
		P = append(P, pair{r[i], l[i]})
	}
	sort.Slice(P, func(i, j int) bool {
		if P[i].x == P[j].x {
			return P[i].y < P[j].y
		}
		return P[i].x < P[j].x
	})

	dp := make([][]int, len(t))
	for i := range dp {
		dp[i] = make([]int, len(t))
	}
	dp[0][0] = 1

	for i := 0; i < M; i++ {
		ndp := make([][]int, len(t))
		for i := range ndp {
			ndp[i] = make([]int, len(t))
		}
		L := P[i].y
		R := P[i].x
		for j := 0; j < len(t); j++ {
			for k := 0; k < len(t); k++ {
				if dp[j][k] == 0 {
					continue
				}
				ndp[j][k] += dp[j][k]
				ndp[j][k] %= mod
				if L > j {
					ndp[j][k] += dp[j][k]
					ndp[j][k] %= mod
				} else if L > k {
					ndp[R][k] += dp[j][k]
					ndp[R][k] %= mod
				} else {
					ndp[R][j] += dp[j][k]
					ndp[R][j] %= mod
				}
			}
		}
		dp, ndp = ndp, dp
	}

	fmt.Println(dp[len(t)-1][len(t)-1])
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
