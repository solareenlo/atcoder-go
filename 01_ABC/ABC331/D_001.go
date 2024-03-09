package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var s [1007][1007]int

func f(x, y int) int {
	return s[n][n]*(x/n)*(y/n) + s[x%n][n]*(y/n) + s[n][y%n]*(x/n) + s[x%n][y%n]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &n, &q)

	for i := 1; i <= n; i++ {
		var c string
		fmt.Fscan(in, &c)
		for j := 1; j <= n; j++ {
			if c[j-1] == 'B' {
				s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + 1
			} else {
				s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1]
			}
		}
	}

	for q > 0 {
		q--
		var xa, ya, xb, yb int
		fmt.Fscan(in, &xa, &ya, &xb, &yb)
		xa++
		ya++
		xb++
		yb++
		fmt.Fprintln(out, f(xb, yb)-f(xa-1, yb)-f(xb, ya-1)+f(xa-1, ya-1))
	}
}
