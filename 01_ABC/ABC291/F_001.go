package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s [100100]string
	var l, r [100100]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		l[i] = n
		r[i] = n
	}
	l[1] = 0
	r[n] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			if i+j+1 <= n && s[i][j] == '1' {
				l[i+j+1] = min(l[i+j+1], l[i]+1)
			}
		}
	}
	for i := n; i >= 1; i-- {
		for j := 0; j < m; j++ {
			if i-j-1 >= 1 && s[i-j-1][j] == '1' {
				r[i-j-1] = min(r[i-j-1], r[i]+1)
			}
		}
	}
	for i := 2; i < n; i++ {
		ans := n
		for j := max(1, i-m+1); j < i; j++ {
			for k := min(n, j+m); k > i; k-- {
				if s[j][k-j-1] == '1' {
					ans = min(ans, l[j]+r[k]+1)
				}
			}
		}
		if ans >= n {
			ans = -1
		}
		fmt.Printf("%d ", ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
