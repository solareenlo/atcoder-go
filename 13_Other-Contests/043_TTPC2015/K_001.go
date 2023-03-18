package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })

	var rest int64
	for i := N - 1; i >= 0; i-- {
		game := int64(i * (i - 1) * (i - 2) / 6)
		if rest+game < A[i] {
			fmt.Println("NO")
			return
		}
		rest += game - A[i]
	}

	fmt.Println("YES")
}
