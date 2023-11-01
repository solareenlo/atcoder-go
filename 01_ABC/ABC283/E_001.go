package main

import (
	"bufio"
	"fmt"
	"os"
)

var m int
var a [1010][1010]int

func check(num, x, y int) bool {
	for i := 1; i <= m; i++ {
		if a[num][i] != a[num][i-1] && a[num][i] != a[num][i+1] && a[num][i] != (a[num-1][i]^x) && a[num][i] != (a[num+1][i]^y) {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [1010][2][2]int

	var n int
	fmt.Fscan(in, &n, &m)
	for i := range a {
		for j := range a[i] {
			a[i][j] = -1
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	for i := 0; i < 1010; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				dp[i][j][k] = 1 << 30
			}
		}
	}
	dp[1][1][0] = 1
	dp[1][0][0] = 0
	for i := 2; i <= n+1; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				for x := 0; x < 2; x++ {
					if check(i-1, k^x, j^k) {
						dp[i][j][k] = min(dp[i][j][k], dp[i-1][k][x]+j)
					}
				}
			}
		}
	}
	ans := min(dp[n+1][0][1], dp[n+1][0][0], dp[n+1][1][1], dp[n+1][1][0])
	if ans > n {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
