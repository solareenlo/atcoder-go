package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 405

var (
	f = [N][N][N]int{}
	n int
	a = [N]int{}
)

func dp(l, r, m int) int {
	if f[l][r][m] != 0 {
		return f[l][r][m]
	}
	t := 0
	if l > 0 || r <= n {
		if m < n-(r-l-1) {
			if a[l] < a[r] {
				t = max(t, dp(l, r+1, m+1))
			} else {
				if l-1 >= 0 {
					t = max(t, dp(l-1, r, m+1))
				}
			}
		}
		if m >= 1 && l >= 1 {
			t = max(t, a[l]+dp(l-1, r, m-1))
		}
		if m >= 1 && r <= n {
			t = max(t, a[r]+dp(l, r+1, m-1))
		}
	}
	f[l][r][m] = t
	return f[l][r][m]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	a[0] = -1
	a[n+1] = -1
	for i := 0; i <= n; i++ {
		fmt.Println(dp(i, i+1, 1))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
