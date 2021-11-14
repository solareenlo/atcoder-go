package main

import "fmt"

func main() {
	dp := [101][101][101]float64{}
	for i := 99; i >= 0; i-- {
		for j := 99; j >= 0; j-- {
			for k := 99; k >= 0; k-- {
				sum := float64(i + j + k)
				if sum == 0.0 {
					continue
				}
				dp[i][j][k] = float64(i)/sum*dp[i+1][j][k] + float64(j)/sum*dp[i][j+1][k] + float64(k)/sum*dp[i][j][k+1] + 1.0
			}
		}
	}

	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(dp[a][b][c])
}
