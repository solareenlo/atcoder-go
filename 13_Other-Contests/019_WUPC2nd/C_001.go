package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	var c, t [30][30]int
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < m; j++ {
			c[i][j] = int(s[j] - '0')
		}
	}
	maxans := 0
	for r := 0; r < 4; r++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				ans := 0
				flag := true
				for k := 0; k < min(n-i, m-j); k++ {
					for l := 0; l < k+1; l++ {
						if c[i+l][j+k-l] == 40 {
							flag = false
						} else {
							ans += c[i+l][j+k-l]
						}
					}
					if flag {
						maxans = max(maxans, ans)
					} else {
						break
					}
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				t[i][j] = c[i][j]
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				c[m-j-1][i] = t[i][j]
			}
		}
		n, m = m, n
	}
	fmt.Println(maxans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
