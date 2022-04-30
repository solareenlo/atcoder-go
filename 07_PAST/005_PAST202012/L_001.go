package main

import "fmt"

func main() {
	var N int
	var S, T string
	fmt.Scan(&N, &S, &T)
	ok := [101][101]bool{}
	for i := 0; i < N; i++ {
		ok[i][i] = true
	}

	for k := 3; k <= N; k += 3 {
		for i := 0; i+k <= N; i++ {
			if S[i] == T[0] && S[i+k-1] == T[2] {
				for j := i + 1; j < i+k-1; j++ {
					if S[j] == T[1] && ok[i+1][j] && ok[j+1][i+k-1] {
						ok[i][i+k] = true
					}
				}
				for j := i + 3; j < i+k; j += 3 {
					if ok[i][j] && ok[j][i+k] {
						ok[i][i+k] = true
					}
				}
			}
		}
	}

	dp := [101]int{}
	for i := 0; i < N; i++ {
		dp[i+1] = max(dp[i+1], dp[i])
		for k := 3; i+k <= N; k += 3 {
			if ok[i][i+k] {
				dp[i+k] = max(dp[i+k], dp[i]+k/3)
			}
		}
	}
	fmt.Println(dp[N])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
