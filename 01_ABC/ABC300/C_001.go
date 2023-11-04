package main

import (
	"fmt"
	"strings"
)

func main() {
	var c [110]string
	var cnt [110]int

	var n, m int
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&c[i])
		c[i] = " " + c[i] + " "
	}
	c[n+1] = strings.Repeat(" ", 110)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if c[i][j] != '#' {
				continue
			}
			d := 1
			for c[i+d][j+d] == '#' && c[i+d][j-d] == '#' && c[i-d][j+d] == '#' && c[i-d][j-d] == '#' {
				d++
			}
			cnt[d-1]++
		}
	}
	for i := 1; i <= min(n, m); i++ {
		fmt.Printf("%d ", cnt[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
