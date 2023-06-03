package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [409][409]int
	var A [409]int

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < 2*N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < 409; i++ {
		for j := 0; j < 409; j++ {
			dp[i][j] = int(1e8)
		}
	}
	for i := 0; i < 2*N; i++ {
		dp[i][i] = 0
	}
	for l := 1; l <= 2*N; l += 2 {
		for i := 0; i+l < 2*N; i++ {
			j := i + l
			dp[i][j+1] = dp[i+1][j] + abs(A[i]-A[j])
			for k := i + 2; k < j; k += 2 {
				dp[i][j+1] = min(dp[i][j+1], dp[i][k]+dp[k][j+1])
			}
		}
	}
	fmt.Println(dp[0][2*N])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
