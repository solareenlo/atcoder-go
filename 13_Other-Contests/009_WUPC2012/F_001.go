package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var x, y [10005]int
	var a [1005][1005]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		x[i]++
		y[i]++
		a[x[i]][y[i]] = 1
	}

	var sum [1005][1005]int
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 1000; j++ {
			sum[i][j] = a[i][j] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if x[i] < x[j] && y[i] < y[j] && a[x[i]][y[j]] == 1 && a[x[j]][y[i]] == 1 && (sum[x[j]-1][y[j]-1]-sum[x[i]][y[j]-1]-sum[x[j]-1][y[i]]+sum[x[i]][y[i]]) == 0 {
				ans = max(ans, (x[i]-x[j])*(y[i]-y[j]))
			}
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
