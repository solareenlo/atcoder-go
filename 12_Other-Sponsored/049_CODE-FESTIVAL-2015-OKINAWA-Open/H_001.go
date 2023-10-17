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

	var N int
	fmt.Fscan(in, &N)
	L := make([]int, N)
	R := make([]int, N)
	X := make([]int, 0)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &L[i], &R[i])
		X = append(X, L[i])
		X = append(X, R[i])
	}
	sort.Ints(X)
	X = unique(X)
	S := len(X)
	G := make([][]int, S+1)
	for i := 0; i < N; i++ {
		L[i] = lowerBound(X, L[i])
		R[i] = lowerBound(X, R[i])
		G[R[i]] = append(G[R[i]], L[i])
	}
	cover := make([][]bool, S-1)
	for i := range cover {
		cover[i] = make([]bool, S)
	}
	for i := 0; i < S-1; i++ {
		cover[i][i] = true
		sum := make([]int, S+1)
		sum[i+1] = 1
		for j := i + 1; j < S; j++ {
			for _, k := range G[j] {
				if k < i {
					continue
				}
				if sum[j] > sum[k] {
					cover[i][j] = true
				}
			}
			if cover[i][j] {
				sum[j+1] = sum[j] + 1
			} else {
				sum[j+1] = sum[j]
			}
		}
	}
	dp0 := make([]int, S)
	dp1 := make([]int, S)
	dp0[0] = 1
	for i := 0; i < S-1; i++ {
		dp0[i+1] += dp0[i]
		if dp0[i+1] >= mod {
			dp0[i+1] -= mod
		}
		dp0[i+1] += dp1[i]
		if dp0[i+1] >= mod {
			dp0[i+1] -= mod
		}
		for j := i + 1; j < S; j++ {
			if cover[i][j] {
				dp1[j] += dp0[i]
				if dp1[j] >= mod {
					dp1[j] -= mod
				}
			}
		}
	}
	fmt.Println((dp0[S-1] + dp1[S-1]) % mod)
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
