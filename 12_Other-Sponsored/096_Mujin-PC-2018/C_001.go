package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 2005

	var w, h [MX][MX]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	var s [MX]string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		s[i] += " "
		t := 0
		for j := 0; j < m+1; j++ {
			if s[i][j] != '.' {
				for k := t; k < j; k++ {
					w[i][k] = j - t - 1
				}
				t = j + 1
			}
		}
	}
	s[n] = strings.Repeat(" ", m+1)
	for j := 0; j < m; j++ {
		t := 0
		for i := 0; i < n+1; i++ {
			if s[i][j] != '.' {
				for k := t; k < i; k++ {
					h[k][j] = i - t - 1
				}
				t = i + 1
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans += w[i][j] * h[i][j]
		}
	}
	fmt.Println(ans)
}
