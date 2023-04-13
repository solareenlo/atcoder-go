package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var f [400][400]int
	for i := range f {
		for j := range f[i] {
			f[i][j] = -int(1e18)
		}
	}
	f[0][0] = 0
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	var a [400]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			for t := max(0, i-k); t < i; t++ {
				f[i][j+1] = max(f[i][j+1], f[t][j]+a[i])
			}
		}
	}
	mx := -int(1e18)
	for i := n - k + 1; i <= n; i++ {
		mx = max(mx, f[i][m])
	}
	if mx < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(mx)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
