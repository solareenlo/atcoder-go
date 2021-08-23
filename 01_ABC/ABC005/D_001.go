package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := [51][51]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&sum[i+1][j+1])
			sum[i+1][j+1] += sum[i+1][j] + sum[i][j+1] - sum[i][j]
		}
	}

	mx := [2501]int{}
	for b := 0; b < n+1; b++ {
		for t := 0; t < b; t++ {
			for r := 0; r < n+1; r++ {
				for l := 0; l < r; l++ {
					mx[(b-t)*(r-l)] = max(mx[(b-t)*(r-l)], sum[b][r]-sum[b][l]-sum[t][r]+sum[t][l])
				}
			}
		}
	}
	for i := 0; i < n*n; i++ {
		mx[i+1] = max(mx[i+1], mx[i])
	}

	var q, p int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&p)
		fmt.Println(mx[p])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
