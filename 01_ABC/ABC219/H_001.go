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

	a := make([]int, n+1)
	x := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &a[i])
	}

	n++
	for j := 0; j < n; j++ {
		for i := 0; i < n-1; i++ {
			if x[i] > x[i+1] {
				t := x[i]
				x[i] = x[i+1]
				x[i+1] = t
				t = a[i]
				a[i] = a[i+1]
				a[i+1] = t
			}
		}
	}

	var start int
	for i := 0; i < n; i++ {
		if x[i] == 0 && a[i] == 0 {
			start = i
		}
	}

	const INF = 1 << 60
	dp := [310][310][2][310]int{}
	for l := 0; l < n; l++ {
		for r := l; r < n+1; r++ {
			for f := 0; f < 2; f++ {
				for c := 0; c < n; c++ {
					dp[l][r][f][c] = -INF
				}
			}
		}
	}
	for c := 0; c < n; c++ {
		dp[start][start+1][0][c] = 0
	}

	for w := 1; w < n; w++ {
		for l := 0; l < n-w+1; l++ {
			r := l + w
			for f := 0; f < 2; f++ {
				for c := 0; c < n; c++ {
					if dp[l][r][f][c] != -INF {
						crr := dp[l][r][f][c]
						var pos int
						if f != 0 {
							pos = x[r-1]
						} else {
							pos = x[l]
						}

						if l != 0 {
							dp[l-1][r][0][c] = max(dp[l-1][r][0][c], crr-(pos-x[l-1])*c)
							if c != 0 {
								dp[l-1][r][0][c-1] = max(dp[l-1][r][0][c-1], crr-(pos-x[l-1])*c+a[l-1])
							}
						}
						if r != n {
							dp[l][r+1][1][c] = max(dp[l][r+1][1][c], crr-(x[r]-pos)*c)
							if c != 0 {
								dp[l][r+1][1][c-1] = max(dp[l][r+1][1][c-1], crr-(x[r]-pos)*c+a[r])
							}
						}
					}
				}
			}
		}
	}

	ans := 0
	for l := 0; l < n; l++ {
		for r := l; r < n+1; r++ {
			for f := 0; f < 2; f++ {
				ans = max(ans, dp[l][r][f][0])
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
