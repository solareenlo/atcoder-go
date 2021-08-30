package main

import "fmt"

func main() {
	var w, n, k, a, b int
	fmt.Scan(&w, &n, &k)

	var dp [10001][51]int = [10001][51]int{}
	for c := 0; c < n; c++ {
		fmt.Scan(&a, &b)
		for i := w; i > a-1; i-- {
			for j := k; j > 0; j-- {
				if dp[i][j] < dp[i-a][j-1]+b {
					dp[i][j] = dp[i-a][j-1] + b
				}
			}
		}
	}
	fmt.Println(dp[w][k])
}
