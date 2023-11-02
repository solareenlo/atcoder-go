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

	var f [505][505]int
	var a [505]int
	B := int(1e14)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		var ch string
		fmt.Fscan(in, &ch)
		for j := 1; j <= n; j++ {
			if ch[j-1] == 'N' {
				f[i][j] = int(1e18)
			} else {
				f[i][j] = B - a[j]
			}
		}
	}
	for i := 1; i <= n; i++ {
		f[i][i] = 0
	}
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j])
			}
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		if f[x][y] >= int(1e17) {
			fmt.Fprintln(out, "Impossible")
		} else {
			p := f[x][y] - a[x]
			fmt.Fprintln(out, (p+B-1)/B, (p+B-1)/B*B-p)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
