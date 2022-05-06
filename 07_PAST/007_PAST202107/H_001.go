package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	SA := 0
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(in, &A)
		SA += A
	}

	f := make([]float64, 101)
	for i := 0; i <= SA; i++ {
		f[i] = math.Hypot(float64(i), 1)
	}

	dp := [101][101][101]float64{}
	for i := 0; i < N; i++ {
		for j := 0; j <= SA; j++ {
			for k := 0; k <= SA; k++ {
				dp[i][j][k] = 1e9
			}
		}
	}

	dp[0][0][SA] = 0
	for i := 0; i < N-1; i++ {
		for j := 0; j <= SA; j++ {
			for k := 0; k <= SA; k++ {
				if dp[i][j][k] < 1e8 {
					for l := 0; l <= k; l++ {
						dp[i+1][l][k-l] = min(dp[i+1][l][k-l], dp[i][j][k]+f[abs(j-l)])
					}
				}
			}
		}
	}
	fmt.Println(dp[N-1][0][0])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
