package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var N int
	fmt.Fscan(in, &N)
	var A [3000]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	var M int
	fmt.Fscan(in, &M)

	var L2RB [3000][]pair
	for i := 0; i < M; i++ {
		var l, r, b int
		fmt.Fscan(in, &l, &r, &b)
		L2RB[l-1] = append(L2RB[l-1], pair{r - 1, b})
	}

	var dp [3001]int
	for i := 0; i < N; i++ {
		sum := -A[i]
		for j := i; j >= 0; j-- {
			for _, rb := range L2RB[j] {
				if rb.x >= i {
					sum += rb.y
				}
			}
			if j != 0 {
				dp[i] = max(dp[i], sum+dp[j-1])
			} else {
				dp[i] = max(dp[i], sum)
			}
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
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
