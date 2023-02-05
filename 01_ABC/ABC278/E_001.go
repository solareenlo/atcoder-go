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

	const N = 303

	var n, m, k, x, y int
	fmt.Fscan(in, &n, &m, &k, &x, &y)

	var a [N][N]int
	var s [N][N][N]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &a[i][j])
			for t := 1; t <= k; t++ {
				s[i][j][t] = s[i-1][j][t] + s[i][j-1][t] - s[i-1][j-1][t]
				if a[i][j] == t {
					s[i][j][t]++
				}
			}
		}
	}

	for i := 0; i <= n-x; i++ {
		for j := 0; j <= m-y; j++ {
			c := 0
			for t := 1; t <= k; t++ {
				if s[i+x][j+y][t]-s[i+x][j][t]-s[i][j+y][t]+s[i][j][t] < s[n][m][t] {
					c++
				}
			}
			fmt.Fprintf(out, "%d ", c)
		}
		fmt.Fprintln(out)
	}
}
