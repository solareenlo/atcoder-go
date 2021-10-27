package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	a := [82][82]int{}
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			var b int
			fmt.Scan(&b)
			a[i][j] = abs(a[i][j] - b)
		}
	}

	dp := [82][82][(80+80)*80 + 100]int{}
	dp[1][1][a[1][1]] = 1

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			for k := 0; k < (80+80)*80; k++ {
				if dp[i][j][k] == 0 {
					continue
				}
				dp[i][j+1][k+a[i][j+1]] = 1
				dp[i][j+1][abs(k-a[i][j+1])] = 1
				dp[i+1][j][k+a[i+1][j]] = 1
				dp[i+1][j][abs(k-a[i+1][j])] = 1
			}
		}
	}

	for k := 0; k < (80+80)*80; k++ {
		if dp[h][w][k] != 0 {
			fmt.Println(k)
			break
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
