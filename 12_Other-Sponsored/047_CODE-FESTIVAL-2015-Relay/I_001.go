package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K, D, a, b int64
	var p []int64
	fmt.Scanf("%d", &N)

	for i := int64(0); i < N; i++ {
		fmt.Scanf("%d %d", &a, &b)
		p = append(p, a+b)
		K += a
		D += b
	}
	if K != D {
		fmt.Println("invalid")
		return
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i] < p[j]
	})

	dp := make([]bool, 40404)
	dp[0] = true

	for _, i := range p {
		for j := 40404 - i - 1; j >= 0; j-- {
			dp[j+i] = dp[j+i] || dp[j]
		}
	}

	if dp[K] {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}
}
