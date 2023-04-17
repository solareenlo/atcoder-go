package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K, X, Y int
	fmt.Fscan(in, &N, &M, &K, &X, &Y)

	var edges [2000][]int
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}

	GCP := "GCP"
	var C [2000]int
	for i := 0; i < N; i++ {
		var c string
		fmt.Fscan(in, &c)
		C[i] = strings.Index(GCP, c)
	}

	var D [2000]int
	for i := 0; i < K; i++ {
		var d int
		fmt.Fscan(in, &d)
		D[i] = C[d-1]
	}
	S := [3]int{Y, 0, X}

	dp := make([]int, N)
	for i := range dp {
		dp[i] = -int(1e9)
	}
	dp[0] = 0
	for t := 0; t < K; t++ {
		dp2 := make([]int, N)
		copy(dp2, dp)
		for i := 0; i < N; i++ {
			for _, j := range edges[i] {
				dp2[i] = max(dp2[i], dp[j])
			}
		}
		for i := 0; i < N; i++ {
			dp[i] = dp2[i] + S[(C[i]-D[t]+3)%3]
		}
	}

	ans := 0
	for i := range dp {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
