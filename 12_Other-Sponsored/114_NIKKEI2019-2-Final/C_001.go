package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var flag [3003][3003]bool
	var a, b [3003][3003]int

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	OK := 0
	for i := 1; i <= k; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		flag[x][y] = true
		OK++
	}
	if OK == n*m {
		fmt.Println(0)
		return
	}
	for j := m; j >= 1; j-- {
		for i := n; i >= 1; i-- {
			if !flag[i][j] {
				a[i][j] = a[i+1][j] + 1
			}
		}
	}
	for i := n; i >= 1; i-- {
		for j := m; j >= 1; j-- {
			if !flag[i][j] {
				b[i][j] = b[i+1][j+1] + 1
			}
		}
	}
	ans := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := ans; i+k <= n && j+k <= m && k <= a[i][j] && k <= b[i][j]; k++ {
				ii := i + k
				jj := j + k
				A := a[i][j] - a[ii+1][j]
				B := a[i][jj] - a[ii+1][jj]
				C := b[i][j] - b[ii+1][jj+1]
				if A == B && B == C && C == k+1 {
					ans = k + 1
				}
			}
		}
	}
	fmt.Println(ans)
}
