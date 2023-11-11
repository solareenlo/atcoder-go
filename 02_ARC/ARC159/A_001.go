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

	const INF = int(1e18)

	var a [105][105]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &a[i][j])
			if a[i][j] == 0 {
				a[i][j] = INF
			}
		}
	}
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				a[i][j] = min(a[i][j], a[i][k]+a[k][j])
			}
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 1; i <= q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x = (x-1)%n + 1
		y = (y-1)%n + 1
		if a[x][y] == INF {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, a[x][y])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
