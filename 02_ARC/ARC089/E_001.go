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

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := [300][300]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	f := [300][300]int{}
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 1; k <= n; k++ {
				for l := 1; l <= m; l++ {
					f[i][j] = max(f[i][j], a[k][l]-i*k-j*l)
				}
			}
		}
	}

	for k := 1; k <= n; k++ {
		for l := 1; l <= m; l++ {
			p := 1 << 60
			for i := 0; i <= 100; i++ {
				for j := 0; j <= 100; j++ {
					p = min(p, f[i][j]+i*k+j*l)
				}
			}
			if p != a[k][l] {
				fmt.Fprintln(out, "Impossible")
				return
			}
		}
	}

	fmt.Fprintln(out, "Possible")
	fmt.Fprintln(out, "202 10401")
	for i := 1; i <= 100; i++ {
		fmt.Fprintln(out, i, i+1, "X")
	}
	for i := 102; i < 202; i++ {
		fmt.Fprintln(out, i, i+1, "Y")
	}
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			fmt.Fprintln(out, i+1, 202-j, f[i][j])
		}
	}
	fmt.Fprintln(out, "1 202")
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
