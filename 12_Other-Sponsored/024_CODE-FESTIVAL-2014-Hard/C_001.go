package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, K int
	fmt.Fscan(in, &n, &m, &K)
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &A[i][j])
			A[i][j]--
		}
	}
	dp := make([]int, (n+1)*(m+1)*K)

	var f func(int, int, int) int
	f = func(i, j, k int) int {
		return (i*(m+1)+j)*K + k
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a := A[i][j]
			for k := 0; k < K; k++ {
				dp[f(i+1, j+1, k)] = dp[f(i, j+1, k)] + dp[f(i+1, j, k)] - dp[f(i, j, k)]
				if k == a {
					dp[f(i+1, j+1, k)]++
				}
			}
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var t, x1, y1, x2, y2 int
		fmt.Fscan(in, &t, &x1, &y1, &x2, &y2)
		if t == 1 {
			x1--
			y1--
			x2--
			y2--
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			a := A[x1][y1]
			b := A[x2][y2]
			A[x1][y1] = b
			A[x2][y2] = a
			if x1 == x2 {
				for i := x1 + 1; i <= n; i++ {
					dp[f(i, y1+1, a)]--
					dp[f(i, y1+1, b)]++
				}
			} else {
				for i := y1 + 1; i <= m; i++ {
					dp[f(x1+1, i, a)]--
					dp[f(x1+1, i, b)]++
				}
			}
		} else {
			ma := -1
			var ans int
			x1--
			y1--
			for k := 0; k < K; k++ {
				x := dp[f(x2, y2, k)] - dp[f(x1, y2, k)] - dp[f(x2, y1, k)] + dp[f(x1, y1, k)]
				if x >= ma {
					ma = x
					ans = k + 1
				}
			}
			fmt.Fprintln(out, ans, ma)
		}
	}
}
