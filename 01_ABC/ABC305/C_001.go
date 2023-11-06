package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	s := make([]string, n+2)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
		s[i] = " " + s[i] + " "
	}
	s[0] = strings.Repeat(" ", 505)
	s[n+1] = strings.Repeat(" ", 505)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			tmp0, tmp1, tmp2, tmp3 := 0, 0, 0, 0
			if s[i+1][j] == '#' {
				tmp0 = 1
			}
			if s[i-1][j] == '#' {
				tmp1 = 1
			}
			if s[i][j+1] == '#' {
				tmp2 = 1
			}
			if s[i][j-1] == '#' {
				tmp3 = 1
			}
			if tmp0+tmp1+tmp2+tmp3 > 1 && s[i][j] != '#' {
				fmt.Println(i, j)
				return
			}
		}
	}
}
