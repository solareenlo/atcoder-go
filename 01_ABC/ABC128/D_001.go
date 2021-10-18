package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, K int
	fmt.Scan(&n, &K)

	v := make([]int, n)
	for i := range v {
		fmt.Scan(&v[i])
	}

	maxi := -(1 << 60)
	for i := 0; i < min(n, K)+1; i++ {
		for j := 0; j < min(i, K-i)+1; j++ {
			for k := 0; k < i; k++ {
				sum := 0
				A := make([]int, 0)
				for l := 0; l < i; l++ {
					A = append(A, v[(l-k+n)%n])
					sum += A[l]
				}
				sort.Ints(A)
				for l := 0; l < j; l++ {
					if A[l] < 0 {
						sum -= A[l]
					}
				}
				maxi = max(maxi, sum)
			}
		}
	}
	fmt.Println(maxi)
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
