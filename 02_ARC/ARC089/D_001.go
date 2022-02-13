package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	x, y, c := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		var C string
		fmt.Scan(&x[i], &y[i], &C)
		x[i] %= k * 2
		y[i] %= k * 2
		if C == "W" {
			c[i] = 1
		}
	}

	white := make([][]int, k)
	for i := range white {
		white[i] = make([]int, k)
	}
	numBlack := 0
	for i := 0; i < n; i++ {
		c[i] += x[i]/k + y[i]/k
		c[i] %= 2
		x[i] %= k
		y[i] %= k
		if c[i] == 1 {
			white[x[i]][y[i]]++
		} else {
			white[x[i]][y[i]]--
		}
		if c[i] == 0 {
			numBlack++
		}
	}

	for i := 0; i < k; i++ {
		for j := 1; j < k; j++ {
			white[i][j] += white[i][j-1]
		}
	}
	for i := 1; i < k; i++ {
		for j := 0; j < k; j++ {
			white[i][j] += white[i-1][j]
		}
	}

	maxi := 0
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			w := white[k-1][k-1] - white[i][k-1] - white[k-1][j] + white[i][j]*2
			w += numBlack
			maxi = max(maxi, max(w, n-w))
		}
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
