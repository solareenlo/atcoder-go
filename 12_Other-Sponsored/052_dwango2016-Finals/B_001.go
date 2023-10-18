package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INF = 1e+10

var x [2][]int
var a, b int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	x[1] = append(x[1], 0)
	for i := 0; i < n; i++ {
		var d int
		fmt.Fscan(in, &d)
		if d < 0 {
			x[0] = append(x[0], -d)
		} else {
			x[1] = append(x[1], d)
		}
	}
	x[0] = append(x[0], 0)
	reverseOrderInt(x[0])
	a = len(x[0])
	b = len(x[1])
	low := 0.0
	up := INF
	for i := 0; i < 100; i++ {
		mid := (low + up) / 2
		if C(mid) {
			up = mid
		} else {
			low = mid
		}
	}
	fmt.Printf("%.12f\n", low)
}

func C(y float64) bool {
	var dp [1010][1010][2]float64
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			dp[i][j][0] = INF * 2
			dp[i][j][1] = INF * 2
		}
	}
	dp[0][0][0] = 0
	dp[0][0][1] = 0
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if dp[i][j][0] > float64(x[0][i]) {
				dp[i][j][0] = INF * 2
			}
			if dp[i][j][1] > float64(x[1][j]) {
				dp[i][j][1] = INF * 2
			}
			if i+1 < a {
				dp[i+1][j][0] = math.Min(dp[i+1][j][0], math.Min(dp[i][j][0]+float64(x[0][i+1]-x[0][i])/y, dp[i][j][1]+(float64)(x[0][i+1]+x[1][j])/y))
			}
			if j+1 < b {
				dp[i][j+1][1] = math.Min(dp[i][j+1][1], math.Min(dp[i][j][0]+(float64)(x[0][i]+x[1][j+1])/y, dp[i][j][1]+(float64)(x[1][j+1]-x[1][j])/y))
			}
		}
	}
	return dp[a-1][b-1][0] <= INF || dp[a-1][b-1][1] <= INF
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
